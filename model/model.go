package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Greeting struct {
	bun.BaseModel `bun:"table:grettings,alias:g"`	
	ID   string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name string `bun:",notnull"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name          string    `bun:"name,notnull"`
	Email         string    `bun:"email,notnull,unique"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

type Article struct {
	bun.BaseModel `bun:"table:articles,alias:a"`
	ID            string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Title         string    `bun:"title,notnull"`
	Thumbnail     string    `bun:"thumbnail,notnull"`
	Content       string    `bun:"content,notnull"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
