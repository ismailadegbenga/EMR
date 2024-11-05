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
	NextOfKinPhone string `json:"nextofkinphone"`
	NextOfKinEmail string `json:"nextofkinemail"`
}

type Staff struct {
	Person
	RoleID uint `json:"role"`
}

type Role struct {
    ID uint `json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt time.Time
    Name string `json: "name"`
}

type InvestigationCategory struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string `json:"name"`
	
}

type Investigation struct {
	ID        uint `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string `json:"name"`
    CategoryID  uint `json:"categoryid"`
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
	{Person: Person{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{},
		Firstname: "Ismail", Middlename: "Abiodun", Lastname: "Adegbenga", Sex: "Male", DOB: time.Date(1982, 9, 18, 0, 0, 0, 0, time.UTC),
		Phone: "08064386157", Email: "ismail@adegbenga.com", Address: "40, Orimolade crescent", Street: "Adeniyi Jones", City: "Ikeja",
		State: "Lagos", Country: "Nigeria"},
		NextOfKinPhone: "08039275448", 
        NextOfKinEmail: "sade.adebanjo@gmail.com",},
	{Person: Person{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{},
    Firstname: "Sulaiman", Middlename: "Adeola", Lastname: "Atolagbe", Sex: "Male", DOB: time.Date(1986, 7, 24, 0, 0, 0, 0, time.UTC),
    Phone: "08029386159", Email: "sulaiman@atolagbe.com", Address: "39", Street: "Sanni Luba", City: "Abule-egba",
    State: "Lagos", Country: "Nigeria"},
    NextOfKinPhone: "08038285338", 
    NextOfKinEmail: "kulikuli.garri@gmail.com",},
	{Person: Person{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{},
    Firstname: "Yusuf", Middlename: "Anjola", Lastname: "Ireti", Sex: "Male", DOB: time.Date(1990, 7, 05, 0, 0, 0, 0, time.UTC),
    Phone: "08054286268", Email: "yusuf.ireti@yahoo.com", Address: "39", Street: "Ita alapo", City: "Ijebu-Ode",
    State: "Ogun", Country: "Nigeria"},
    NextOfKinPhone: "08059476478", 
    NextOfKinEmail: "rice.beans@gmail.com",},
}

var staff = []Staff{
	{Person: Person{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{},
    Firstname: "Ismail", Middlename: "Abiodun", Lastname: "Adegbenga", Sex: "Male", DOB: time.Date(1982, 9, 18, 0, 0, 0, 0, time.UTC),
    Phone: "08064386157", Email: "ismail@adegbenga.com", Address: "40, Orimolade crescent", Street: "Adeniyi Jones", City: "Ikeja",
    State: "Lagos", Country: "Nigeria"},
    RoleID: 1,,},
	{Person: Person{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{},
    Firstname: "Laduke", Middlename: "Adenrele", Lastname: "Mukaiba", Sex: "Female", DOB: time.Date(19, 9, 18, 0, 0, 0, 0, time.UTC),
    Phone: "08065386459", Email: "laduke@mukaiba.com", Address: "11, Olaide", Street: "Olakitan revue", City: "Gwagwalada",
    State: "Abuja", Country: "Nigeria"},
    RoleID: 2},
}

var invcat = []InvestigationCategory{
	{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}, Name: Radiology},
	{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}, Name: Laboratory},
	{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: time.Time{}, Name: Cardiology},
}

var investigations = []Investigation{
	{ID: 1, CreatedAt: time.Now(), UpdatedAt time.Now(), DeletedAt time.Time{},
	Name: "Xray", CategoryID: 1, Price: 1700000,},
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
