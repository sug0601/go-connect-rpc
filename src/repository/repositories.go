package repository

import "github.com/uptrace/bun"

type Repositories struct {
	Greeting GreetingRepository
	User     UserRepository
	Article  ArticleRepository
}

func NewRepositories(db *bun.DB) *Repositories {
	return &Repositories{
		Greeting: NewBunGreetingRepository(db),
		User:     NewBunUserRepository(db),
		Article:  NewBunArticleRepository(db),
	}
}
