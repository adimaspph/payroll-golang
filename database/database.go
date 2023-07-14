package database

import (
	"encoding/json"
	"github.com/adimaspph/payroll-golang/models"
	"log"
	"os"
)

func ReadCompany() (*models.Company, error) {
	// Open the JSON file
	file, err := os.ReadFile("database/database.json")
	if err != nil {
		log.Println("Error opening file:", err)
		return nil, err
	}

	// Decode the JSON data into a struct
	var company models.Company
	err = json.Unmarshal(file, &company)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, err
	}

	return &company, nil
}

func WriteCompany(company *models.Company) error {
	// Convert struct to JSON
	jsonData, err := json.Marshal(company)
	if err != nil {
		log.Println("Error marshaling JSON:", err)
		return err
	}

	// Write JSON data to a file
	file, err := os.Create("database/database.json")
	if err != nil {
		log.Println("Error creating file:", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Println("Error writing JSON to file:", err)
	}

	log.Println("Person data written to person.json")

	return nil
}
