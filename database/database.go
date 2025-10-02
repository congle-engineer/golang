package database

import (
  "database/sql"
  "fmt"
  "log"
  "strings"

  _ "github.com/lib/pq"
)

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

func Connect(host string, port int, user, password, dbname string) (*sql.DB, error) {
  psqlInfo := fmt.Sprintf(
    "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    return nil, err
  }

  if err := db.Ping(); err != nil {
    return nil, err
  }

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
  if _, err := db.Exec(createTable); err != nil {
    return nil, err
  }

  return db, nil
}

func InsertDelegations(db *sql.DB, delegations []Delegation) error {
  if len(delegations) == 0 {
    return nil
  }

  insertStmt := `
  INSERT INTO delegations 
  (delegator, recipient, delegated_balance, kind, amount, fee_rate, fee, transaction_fee, note)
  VALUES `

  values := []interface{}{}
  placeholders := []string{}

  for i, d := range delegations {
    n := i * 9
    placeholders = append(placeholders,
      fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
        n+1, n+2, n+3, n+4, n+5, n+6, n+7, n+8, n+9))

    values = append(values,
      d.Delegator, d.Recipient, d.DelegatedBalance,
      d.Kind, d.Amount, d.FeeRate, d.Fee, d.TransactionFee, d.Note)
  }

  insertStmt += strings.Join(placeholders, ",")
  insertStmt += " RETURNING id"

  rows, err := db.Query(insertStmt, values...)
  if err != nil {
    return err
  }
  defer rows.Close()

  for rows.Next() {
    var id int
    if err := rows.Scan(&id); err != nil {
      return err
    }
    log.Println("Inserted row with ID:", id)
  }
  return nil
}
