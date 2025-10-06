package repository

import (
	"context"

	"example.com/model"
	"github.com/uptrace/bun"
)

type ArticleRepository interface {
	FindAll(ctx context.Context) ([]*model.Article, error)
	FindByID(ctx context.Context, id string) (*model.Article, error)
	Insert(ctx context.Context, article *model.Article) error
}

type BunArticleRepository struct {
	db *bun.DB
}

func NewBunArticleRepository(db *bun.DB) *BunArticleRepository {
	return &BunArticleRepository{db: db}
}

func (r *BunArticleRepository) FindAll(ctx context.Context) ([]*model.Article, error) {
	var articles []*model.Article
	err := r.db.NewSelect().
		Model(&articles).
		Order("created_at DESC").
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *BunArticleRepository) FindByID(ctx context.Context, id string) (*model.Article, error) {
	article := new(model.Article)
	err := r.db.NewSelect().
		Model(article).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (r *BunArticleRepository) Insert(ctx context.Context, article *model.Article) error {
	_, err := r.db.NewInsert().
		Model(article).
		Exec(ctx)
	return err
}
