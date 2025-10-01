package db

type Greeting struct {
	ID   int64  `bun:",pk,autoincrement"`
	Name string `bun:",notnull"`
}
