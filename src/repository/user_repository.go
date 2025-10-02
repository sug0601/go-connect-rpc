package repository

import (
	"context"

	"example.com/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	Insert(ctx context.Context, name, email string) (*model.User, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type BunUserRepository struct {
	db *bun.DB
}

func NewBunUserRepository(db *bun.DB) *BunUserRepository {
	return &BunUserRepository{db: db}
}

func (r *BunUserRepository) Insert(ctx context.Context, name, email string) (*model.User, error) {
	user := &model.User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}
	_, err := r.db.NewInsert().
		Model(user).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *BunUserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	err := r.db.NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *BunUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	err := r.db.NewSelect().
		Model(user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
