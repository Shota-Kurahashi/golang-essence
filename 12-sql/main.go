package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=example dbname=todo sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	id := 3

	rows, err := db.Query("SELECT * FROM todos WHERE id = $1", id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var age int

		err = rows.Scan(&name, &age)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(name, age)

	}

	var name string
	var age int

	err = db.QueryRow("SELECT * FROM todos WHERE id = $1", id).Scan(&name, &age)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(name, age)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	rows, err = db.QueryContext(ctx, "SELECT * FROM todos WHERE id = $1", id)

	if err != nil {
		log.Fatal(err)
	}

	// Exec
	r, err := db.Exec("INSERT INTO todos (name, age) VALUES ($1, $2)", "John", 20)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)

	stmt, err := db.Prepare("INSERT INTO todos (name, age) VALUES ($1, $2)")

	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow("John", 20).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(id)

	// Transaction
	transactions(db)
}

func transactions(db *sql.DB) error {
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	result, err := tx.Exec("INSERT INTO todos (name, age) VALUES ($1, $2)", "John", 20)

	if err != nil {
		return err
	}

	if affected, err := result.RowsAffected(); err == nil {
		fmt.Println(affected)
	} else if affected == 0 {
		fmt.Println("No rows affected")
	}

	err = tx.Commit()

	if err != nil {
		return err
	}

	return nil

}
