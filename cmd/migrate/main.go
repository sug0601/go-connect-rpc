package main

import (
	"context"
	"log"

	"example.com/model"
	"example.com/src/infra"
	"github.com/uptrace/bun"
)

func main() {
	dbConn := infra.NewPostgresDB("postgres://user:pass@localhost:5433/connect?sslmode=disable")
	defer dbConn.Close()

	ctx := context.Background()

	// 追加する
	models := []interface{}{
		(*model.Greeting)(nil),
		(*model.User)(nil),
	}

	if err := createTables(ctx, dbConn, models); err != nil {
		log.Fatal(err)
	}

	log.Println("All tables created successfully!")
}

func createTables(ctx context.Context, db *bun.DB, models []interface{}) error {
	for _, model := range models {
		_, err := db.NewCreateTable().
			Model(model).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			return err
		}
		log.Printf("Table for %T created or already exists", model)
	}
	return nil
}
