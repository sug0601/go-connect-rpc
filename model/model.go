package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Greeting struct {
	ID   int64  `bun:",pk,autoincrement"`
	Name string `bun:",notnull"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            string    `bun:"id,pk,type:uuid"`
	Name          string    `bun:"name,notnull"`
	Email         string    `bun:"email,notnull,unique"`
	CreatedAt     time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}
