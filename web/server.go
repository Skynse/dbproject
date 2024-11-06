package web

import (
	"database/sql"
	"dbproject/db_core"
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	db *db_core.DBService
}

func NewServer(db *db_core.DBService) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) SetupRoutes() {
	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// API routes
	http.HandleFunc("/api/items", s.getAllItems)      // GET all items
	http.HandleFunc("/api/item/search", s.getItem)    // GET single item by ID or name
	http.HandleFunc("/api/item/insert", s.insertItem) // POST new item
	http.HandleFunc("/api/item", s.modifyItem)        // PUT update and DELETE single item
}

// Fetch single item by ID or name
func (s *Server) getItem(w http.ResponseWriter, r *http.Request) {
	s.enableCors(&w)
	query := r.URL.Query()

	var rows *sql.Rows
	var err error

	if query.Get("name") != "" {
		rows, err = s.db.ExecuteQuery("SELECT * FROM ITEM WHERE Iname LIKE ?", query.Get("name"))
	} else if query.Get("id") != "" {
		rows, err = s.db.ExecuteQuery("SELECT * FROM ITEM WHERE Iid = ?", query.Get("id"))
	} else {
		http.Error(w, "invalid query parameter", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "failed to get item", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var item db_core.Item
	if rows.Next() {
		if err := rows.Scan(&item.IIId, &item.Iname, &item.Sprice, &item.Idescription); err != nil {
			http.Error(w, "failed to scan item", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

// Fetch all items
func (s *Server) getAllItems(w http.ResponseWriter, r *http.Request) {
	s.enableCors(&w)
	rows, err := s.db.ExecuteQuery("SELECT * FROM ITEM")
	if err != nil {
		http.Error(w, "failed to get items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []db_core.Item
	for rows.Next() {
		var item db_core.Item
		if err := rows.Scan(&item.IIId, &item.Iname, &item.Sprice, &item.Idescription); err != nil {
			http.Error(w, "failed to scan items", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Insert new item
func (s *Server) insertItem(w http.ResponseWriter, r *http.Request) {
	s.enableCors(&w)

	// Check for method POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Define the struct to capture the incoming fields (no ID field here)
	var newItem struct {
		Iname        string  `json:"Iname"`
		Sprice       float64 `json:"Sprice"`
		Idescription string  `json:"Idescription"`
	}

	// Decode the request body into newItem
	err := json.NewDecoder(r.Body).Decode(&newItem)
	// if err != nil {
	// 	http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// Debug: Print the decoded item to check if it's properly decoded
	fmt.Printf("Decoded item: %+v\n", newItem)

	// Execute the database query to insert the new item
	_, err = s.db.ExecuteQuery("INSERT INTO ITEM (Iname, Sprice, Idescription) VALUES (?, ?, ?)", newItem.Iname, newItem.Sprice, newItem.Idescription)
	if err != nil {
		http.Error(w, "failed to insert item: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Item inserted successfully")
}

// Update or delete item based on HTTP method
func (s *Server) modifyItem(w http.ResponseWriter, r *http.Request) {
	s.enableCors(&w)
	if r.Method != http.MethodPut && r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var item db_core.Item
		json.NewDecoder(r.Body).Decode(&item)
		// if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		// 	http.Error(w, "invalid request body", http.StatusBadRequest)
		// 	return
		// }
		fmt.Printf("Decoded item: %+v\n", item)
		_, err := s.db.ExecuteQuery("UPDATE ITEM SET Iname = ?, Sprice = ?, Idescription = ? WHERE Iid = ?", item.Iname, item.Sprice, item.Idescription, item.IIId)
		if err != nil {
			http.Error(w, "failed to update item", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Item updated successfully")

	case http.MethodDelete:
		var item struct {
			IIId int `json:"Iid"`
		}
		json.NewDecoder(r.Body).Decode(&item)
		_, err := s.db.ExecuteQuery("DELETE FROM ITEM WHERE Iid = ?", item.IIId)
		if err != nil {
			http.Error(w, "failed to delete item", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "Item deleted successfully")
	}
}

func (s *Server) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8000", nil)
}
