package main

import (
	"fmt"
	"log"

	"golang/database"
)

func main() {
	db, err := database.Connect("localhost", 5432, "postgres", "postgres", "trd")
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()
	fmt.Println("Connected to DB successfully")

	delegations := []database.Delegation{
		{"tz1Alice", "tz1Validator", 1000.0, "delegation", 1000.0, 0.05, 50.0, 0.1, "First"},
		{"tz1Bob", "tz1Validator", 2000.0, "delegation", 2000.0, 0.05, 100.0, 0.2, "Second"},
		{"tz1Charlie", "tz1Validator", 1500.0, "delegation", 1500.0, 0.05, 75.0, 0.15, "Third"},
	}

	if err := database.InsertDelegations(db, delegations); err != nil {
		log.Fatal("Insert error:", err)
	}
}
