package main

import (
	"time"
)

// Data structures:
// Patient
// Staff
// InvestigationCategory
// Investigation
// Request
// Result

type Person struct {
	ID         uint `json:"id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	Firstname  string    `json:"firstname"`
	Middlename string    `json:"middlename"`
	Lastname   string    `json:"lastname"`
	Sex        string    `json:"sex"`
	DOB        time.Time `json:"dob"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Address    string    `json:"address"`
	Street     string    `json:"street"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Country    string    `json:"country"`
}

type Client struct {
	Person
	NextOfKinPhone string `json:"nkphone"`
	NextOfKinEmail string `json:"nkemail"`
}

type Staff struct {
	Person
	Role string `json:"role"`
}

type InvestigationCategory struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string `json:"name"`
	Modality  string `json:"modality"`
}

type Investigation struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string `json:"name"`
	Price     uint   `json:"price"`
}

type Request struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	PatientID uint `json:"patientid"`
	CreatedBy uint `json:"createdby"`
}

type Result struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	RID       uint
	Report    string `json:"report"`
	CreatedBy uint   `json:"createdby"`
}

// Seed structs with data

var clients = []Client{
	{Person{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{},
		Firstname: "Ismail", Middlename: "Abiodun", Lastname: "Adegbenga", Sex: "Male", DOB: time.Time{Year: 1982, Month: 9, Day: 18},
		Phone: "08064386157", Email: "ismail@adegbenga.com", Address: "40, Orimolade crescent", Street: "Adeniyi Jones", City: "Ikeja",
		State: "Lagos", Country: "Nigeria"},
		NextOfKinPhone: "08039275448", NextOfKinEmail: "sade.adebanjo@gmail.com"},
	{},
	{},
}

var staff = []Staff{
	{},
	{},
}

var invcat = []InvestigationCategory{
	{},
	{},
	{},
}

var investigations = []Investigation{
	{},
	{},
	{},
}

var rq = []Request{
	{},
	{},
	{},
}

var rs = []Result{
	{},
	{},
	{},
}
