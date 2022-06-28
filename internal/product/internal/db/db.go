package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

const dsn = "host=localhost port=55000 user=postgres password=postgrespw dbname=hm3 sslmode=disable"

func New(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, dsn)
}
