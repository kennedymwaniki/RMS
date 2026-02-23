package db

import(
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/google/uuid"
)

type pg struct {
	db *pgxpool.Pool
}