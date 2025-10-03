package repository

import (
	"context"

	"example.com/model"
	"github.com/uptrace/bun"
)

type GreetingRepository interface {
	Insert(ctx context.Context, name string) error
	Exists(ctx context.Context, name string) (bool, error)
}

type BunGreetingRepository struct {
	db *bun.DB
}

func NewBunGreetingRepository(db *bun.DB) *BunGreetingRepository {
	return &BunGreetingRepository{db: db}
}

func (r *BunGreetingRepository) Insert(ctx context.Context, name string) error {
	_, err := r.db.NewInsert().
		Model(&model.Greeting{Name: name}).
		Exec(ctx)
	return err
}

func (r *BunGreetingRepository) Exists(ctx context.Context, name string) (bool, error) {
	exists, err := r.db.NewSelect().
		Model((*model.Greeting)(nil)).
		Where("name = ?", name).
		Exists(ctx)
	if err != nil {
		return false, err
	}
	return exists, nil
}
