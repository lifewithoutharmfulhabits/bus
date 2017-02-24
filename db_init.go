package main

import (
	"log"
)

const DDL_BUSES = `
CREATE TABLE IF NOT EXISTS buses (
	bus_id integer PRIMARY KEY AUTOINCREMENT,
	number TEXT UNIQUE NOT NULL,
	description TEXT NULL
)`

const DDL_DRIVERS = `
CREATE TABLE IF NOT EXISTS drivers (
	driver_id integer PRIMARY KEY AUTOINCREMENT,
	first_name TEXT,
	last_name TEXT
)
`
const DDL_PAYMENTS = `
CREATE TABLE IF NOT EXISTS payments(
	payment_id integer PRIMARY KEY AUTOINCREMENT,
	driver_id INTEGER NOT NULL,
	bus_id INTEGER NOT NULL,
	amount REAL,
	comment TEXT
)
`

func createDbIfNotExists() {

	database, operErr := OpenDbConnection()
	if operErr != nil {
		log.Println("Failed to create the handle")
	}
	defer database.Close()

	if pingErr := database.Ping(); pingErr != nil {
		log.Fatal("Failed to keep connection alive")
	}

	if _, err := database.Exec(DDL_BUSES); err != nil {
		log.Fatal(err)
	}
	if _, err := database.Exec(DDL_DRIVERS); err != nil {
		log.Fatal(err)
	}
	if _, err := database.Exec(DDL_PAYMENTS); err != nil {
		log.Fatal(err)
	}
}
