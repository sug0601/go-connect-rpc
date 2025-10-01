package infra

import (
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPostgresDB(dsn string) *bun.DB {
	connector := pgdriver.NewConnector(pgdriver.WithDSN(dsn))

	sqlDB := sql.OpenDB(connector)
	if sqlDB == nil {
		log.Fatal("sql.OpenDB returned nil")
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	db := bun.NewDB(sqlDB, pgdialect.New())
	return db
}
