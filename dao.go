package main

import (
	"log"
)

type Dao struct {
}

var DAO Dao = Dao{}

func (*Dao) addDriver(driver *Driver) error {
	_, err := ExecuteInsert("INSERT INTO drivers (first_name, last_name) VALUES(?, ?)",
		driver.FirstName, driver.LastName)
	return err
}

func (*Dao) getAllDrivers() ([]Driver, error) {
	rows, err := ExecuteSelect("SELECT driver_id, first_name, last_name FROM drivers", nil)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	result := make([]Driver, 0, 3)
	for rows.Next() {
		driver := Driver{}
		err := rows.Scan(&driver.Id, &driver.FirstName, &driver.LastName)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, driver)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}
