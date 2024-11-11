package db_core

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DBService struct {
	DB *sql.DB
}

// NewDBService creates and initializes a new DBService with connection pool settings
func NewDBService() (*DBService, error) {
	db, err := connect()
	if err != nil {
		return nil, fmt.Errorf("failed to create DBService: %w", err)
	}

	return &DBService{
		DB: db,
	}, nil
}

// Close gracefully closes the database connection
func (d *DBService) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}

func connect() (*sql.DB, error) {
	// Get connection details from environment or use defaults
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost" // Default to container service name
	}

	fmt.Printf("Attempting to connect to MySQL at host: %s\n", host)

	dsn := fmt.Sprintf("root:@tcp(%s:3306)/BEST_PRICE_DB?parseTime=true", host)

	var db *sql.DB
	var err error

	// Add retry logic for container startup timing
	for i := 0; i < 30; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Printf("Failed to open database connection: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			fmt.Println("Successfully connected to database!")
			break
		}

		fmt.Printf("Attempt %d: Database ping failed: %v\n", i+1, err)
		db.Close()
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("error connecting to the database after retries: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

// Example of adding a method to execute queries
func (d *DBService) ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error) {
	if d.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	return d.DB.Query(query, args...)
}

// Example of adding a method to execute a single-row query
func (d *DBService) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRow(query, args...)
}

// Example of adding a method to execute statements that don't return rows
func (d *DBService) Execute(query string, args ...interface{}) (sql.Result, error) {
	if d.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	return d.DB.Exec(query, args...)
}

func (d *DBService) GetAllItems() ([]Item, error) {
	rows, err := d.ExecuteQuery("SELECT * FROM ITEM")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.IIId, &item.Iname, &item.Sprice)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}
