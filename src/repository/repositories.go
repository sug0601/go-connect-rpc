package repository

import "github.com/uptrace/bun"

type Repositories struct {
	Greeting GreetingRepository
	User     UserRepository
}

func NewRepositories(db *bun.DB) *Repositories {
	return &Repositories{
		Greeting: NewBunGreetingRepository(db),
		User:     NewBunUserRepository(db),
	}
}
