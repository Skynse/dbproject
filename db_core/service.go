package db_core

import (
	"database/sql"
	"fmt"
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
	dsn := "root:@tcp(localhost:3306)/BEST_PRICE_DB?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify the connection
	err = db.Ping()
	if err != nil {
		db.Close() // Close the database if ping fails
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

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
