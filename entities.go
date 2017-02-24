package main

type Driver struct {
	Id        int
	FirstName string
	LastName  string
}

type Bus struct {
	Id          int
	Number      string
	Description string
}

type Payment struct {
	Id       int
	DriverId int
	BusId    int
	Amount   float32
	Comment  string
}
