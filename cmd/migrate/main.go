package main

import (
	"context"
	"log"

	"example.com/src/db"
	"example.com/src/infra"
)

func main() {
	dbConn := infra.NewPostgresDB("postgres://user:pass@localhost:5433/connect?sslmode=disable")
	defer dbConn.Close()

	ctx := context.Background()
	_, err := dbConn.NewCreateTable().Model((*db.Greeting)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Greeting table created!")
}
