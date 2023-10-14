package main

import (
	"context"
	"ent/ent"
	"log"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres password=example dbname=todo sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	err = client.Schema.Create(context.Background())

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	for _, todo := range client.Todo.Query().AllX(context.Background()) {
		log.Println(todo)
	}

	_, err = client.Todo.Create().SetText("Buy milk").SetPriority(1).Save(context.Background())

	if err != nil {
		log.Fatalln(err)
	}
}
