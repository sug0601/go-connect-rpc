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

	// Truncate tables
	if err := truncateTables(ctx, dbConn); err != nil {
		log.Fatal(err)
	}
	log.Println("All tables truncated successfully!")

	// Seed data
	if err := seed(ctx, dbConn); err != nil {
		log.Fatal(err)
	}
	log.Println("Seed data inserted successfully!")
}

func truncateTables(ctx context.Context, db *bun.DB) error {
	tables := []interface{}{
		(*model.Greeting)(nil),
		(*model.User)(nil),
		(*model.Article)(nil),
	}

	for _, table := range tables {
		_, err := db.NewTruncateTable().
			Model(table).
			Cascade(). // 外部キーがある場合も削除
			Exec(ctx)
		if err != nil {
			return err
		}
		log.Printf("Table %T truncated", table)
	}

	return nil
}

func seed(ctx context.Context, db *bun.DB) error {
	// Greeting
	_, err := db.NewInsert().Model(&model.Greeting{
		Name: "Hello",
	}).Exec(ctx)
	if err != nil {
		return err
	}

	// User
	user := &model.User{
		Name:      "Alice",
		Email:     "alice@example.com",
	}
	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return err
	}

	// Articles
	articles := []*model.Article{
		{
			Title:     "First Article",
			Thumbnail: "https://picsum.photos/seed/1/400/300",
			Content:   "This is the first seeded article.",
		},
		{
			Title:     "Second Article",
			Thumbnail: "https://picsum.photos/seed/2/400/300",
			Content:   "This is the second seeded article.",
		},
		{
			Title:     "Third Article",
			Thumbnail: "https://picsum.photos/seed/3/400/300",
			Content:   "This is the third seeded article.",
		},
		{
			Title:     "Fourth Article",
			Thumbnail: "https://picsum.photos/seed/4/400/300",
			Content:   "This is the fourth seeded article.",
		},
		{
			Title:     "Fifth Article",
			Thumbnail: "https://picsum.photos/seed/5/400/300",
			Content:   "This is the fifth seeded article.",
		},
		{
			Title:     "Sixth Article",
			Thumbnail: "https://picsum.photos/seed/6/400/300",
			Content:   "This is the sixth seeded article.",
		},
		{
			Title:     "Seventh Article",
			Thumbnail: "https://picsum.photos/seed/7/400/300",
			Content:   "This is the seventh seeded article.",
		},
		{
			Title:     "Eighth Article",
			Thumbnail: "https://picsum.photos/seed/8/400/300",
			Content:   "This is the eighth seeded article.",
		},
		{
			Title:     "Ninth Article",
			Thumbnail: "https://picsum.photos/seed/9/400/300",
			Content:   "This is the ninth seeded article.",
		},
		{
			Title:     "Tenth Article",
			Thumbnail: "https://picsum.photos/seed/10/400/300",
			Content:   "This is the tenth seeded article.",
		},
	}


	_, err = db.NewInsert().Model(&articles).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
