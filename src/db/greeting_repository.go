package db

import (
	"context"

	"github.com/uptrace/bun"
)

type GreetingRepository interface {
	Insert(ctx context.Context, name string) error
}

type BunGreetingRepository struct {
	db *bun.DB
}

func NewBunGreetingRepository(db *bun.DB) *BunGreetingRepository {
	return &BunGreetingRepository{db: db}
}

func (r *BunGreetingRepository) Insert(ctx context.Context, name string) error {
	_, err := r.db.NewInsert().
		Model(&Greeting{Name: name}).
		Exec(ctx)
	return err
}
