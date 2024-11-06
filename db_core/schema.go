package db_core

import (
	"time"
)

// Store represents the STORE table.
type Store struct {
	SId     int    `json:"sId"`
	Sname   string `json:"Sname"`
	Street  string `json:"Street,omitempty"`
	City    string `json:"City,omitempty"`
	StateAb string `json:"StateAb,omitempty"`
	ZipCode string `json:"ZipCode,omitempty"`
	Sdate   string `json:"Sdate,omitempty"`
	Telno   string `json:"Telno,omitempty"`
	URL     string `json:"URL,omitempty"`
}

// Customer represents the CUSTOMER table.
type Customer struct {
	CId     int    `json:"cId"`
	Cname   string `json:"Cname"`
	Street  string `json:"Street,omitempty"`
	City    string `json:"City,omitempty"`
	StateAb string `json:"StateAb,omitempty"`
	ZipCode string `json:"ZipCode,omitempty"`
	CTelNo  string `json:"CTelNo,omitempty"`
	Cdob    string `json:"Cdob,omitempty"`
}

// Vendor represents the VENDOR table.
type Vendor struct {
	VId     int    `json:"vId"`
	Vname   string `json:"Vname"`
	Street  string `json:"Street,omitempty"`
	City    string `json:"City,omitempty"`
	StateAb string `json:"StateAb,omitempty"`
	ZipCode string `json:"ZipCode,omitempty"`
	Vemail  string `json:"Vemail,omitempty"`
	VTelNo  string `json:"VTelNo,omitempty"`
}

// VendorStore represents the VENDOR_STORE many-to-many relationship.
type VendorStore struct {
	VId int `json:"vId"`
	SId int `json:"sId"`
}

// Contract represents the CONTRACT table.
type Contract struct {
	VId    int       `json:"vId"`
	CtId   int       `json:"ctId"`
	Sdate  time.Time `json:"Sdate"`
	Ctime  time.Time `json:"Ctime"`
	Ctname string    `json:"Ctname"`
}

// Item represents the ITEM table.
type Item struct {
	IIId         int     `json:"iId"`
	Iname        string  `json:"Iname"`
	Sprice       float64 `json:"Sprice,omitempty"`
	Idescription string  `json:"Idescription,omitempty"`
}

// VendorItem represents the VENDOR_ITEM many-to-many relationship.
type VendorItem struct {
	VId    int     `json:"vId"`
	IIId   int     `json:"iId"`
	Vprice float64 `json:"Vprice,omitempty"`
}

// Order represents the ORDERS table.
type Order struct {
	OId    int       `json:"oId"`
	SId    int       `json:"sId"`
	CId    int       `json:"cId"`
	Odate  time.Time `json:"Odate"`
	Ddate  time.Time `json:"Ddate,omitempty"`
	Amount float64   `json:"Amount,omitempty"`
}

// OrderItem represents the ORDER_ITEM many-to-many relationship.
type OrderItem struct {
	OId    int `json:"oId"`
	SId    int `json:"sId"`
	IIId   int `json:"iId"`
	Icount int `json:"Icount"`
}

// StoreItem represents the STORE_ITEM table.
type StoreItem struct {
	SId    int `json:"sId"`
	IIId   int `json:"iId"`
	Scount int `json:"Scount"`
}

// StoreCustomer represents the STORE_CUSTOMER many-to-many relationship.
type StoreCustomer struct {
	SId int `json:"sId"`
	CId int `json:"cId"`
}

// Employee represents the EMPLOYEE table.
type Employee struct {
	SId       int       `json:"sId"`
	SSN       string    `json:"SSN"`
	Ename     string    `json:"Ename"`
	Street    string    `json:"Street,omitempty"`
	City      string    `json:"City,omitempty"`
	StateAb   string    `json:"StateAb,omitempty"`
	ZipCode   string    `json:"ZipCode,omitempty"`
	Etype     string    `json:"Etype,omitempty"`
	Bdate     time.Time `json:"Bdate,omitempty"`
	Sdate     time.Time `json:"Sdate,omitempty"`
	Edate     time.Time `json:"Edate,omitempty"`
	ELevel    string    `json:"ELevel,omitempty"`
	Asalary   float64   `json:"Asalary,omitempty"`
	Agency    string    `json:"Agency,omitempty"`
	Hsalary   float64   `json:"Hsalary,omitempty"`
	Institute string    `json:"Institute,omitempty"`
	Itype     string    `json:"Itype,omitempty"`
	TelNo     string    `json:"TelNo,omitempty"`
	Email     string    `json:"Email,omitempty"`
}

// Feedback represents the FEEDBACK table.
type Feedback struct {
	CId      int    `json:"cId"`
	RId      int    `json:"rId"`
	Rating   int    `json:"Rating"`
	Comments string `json:"Comments,omitempty"`
}
