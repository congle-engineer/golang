package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

// cấu hình Postgres
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "trd"
)

// Struct mapping với bảng
type Delegation struct {
	Delegator        string
	Recipient        string
	DelegatedBalance float64
	Kind             string
	Amount           float64
	FeeRate          float64
	Fee              float64
	TransactionFee   float64
	Note             string
}

func main() {
	// chuỗi kết nối
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// mở kết nối
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// test kết nối
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}
	fmt.Println("Connected to DB successfully")

	// tạo bảng nếu chưa có
	createTable := `
	CREATE TABLE IF NOT EXISTS delegations (
		id SERIAL PRIMARY KEY,
		delegator TEXT,
		recipient TEXT,
		delegated_balance NUMERIC,
		kind TEXT,
		amount NUMERIC,
		fee_rate NUMERIC,
		fee NUMERIC,
		transaction_fee NUMERIC,
		note TEXT
	)`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Create table error:", err)
	}

	// dữ liệu cần insert (nhiều record)
	delegations := []Delegation{
		{"tz1Alice", "tz1Validator", 1000.0, "delegation", 1000.0, 0.05, 50.0, 0.1, "First"},
		{"tz1Bob", "tz1Validator", 2000.0, "delegation", 2000.0, 0.05, 100.0, 0.2, "Second"},
		{"tz1Charlie", "tz1Validator", 1500.0, "delegation", 1500.0, 0.05, 75.0, 0.15, "Third"},
	}

	// build câu lệnh insert nhiều record
	insertStmt := `
	INSERT INTO delegations 
	(delegator, recipient, delegated_balance, kind, amount, fee_rate, fee, transaction_fee, note)
	VALUES `

	values := []interface{}{}
	placeholders := []string{}

	for i, d := range delegations {
		fmt.Println("Processing record:", d)
		fmt.Println("Index:", i)
		n := i * 9 // mỗi record có 9 field
		placeholders = append(placeholders,
			fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
				n+1, n+2, n+3, n+4, n+5, n+6, n+7, n+8, n+9))
		fmt.Println("Placeholders:", placeholders)

		values = append(values,
			d.Delegator, d.Recipient, d.DelegatedBalance,
			d.Kind, d.Amount, d.FeeRate, d.Fee, d.TransactionFee, d.Note)
		fmt.Println("Values:", values)
	}

	insertStmt += strings.Join(placeholders, ",")
	fmt.Println("Insert statement:", insertStmt)
	insertStmt += " RETURNING id"
	fmt.Println("Insert statement with RETURNING id:", insertStmt)

	rows, err := db.Query(insertStmt, values...)
	if err != nil {
		log.Fatal("Batch insert error:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted row with ID:", id)
	}
}
