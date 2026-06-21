package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"wallet/internal/config"
)

type Storage struct {
	pool *pgxpool.Pool
}

// New Создание нового подключения к бд
func New(db config.DB) (*Storage, error) {
	const op = "storage.New"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode)

	pool, err := pgxpool.New(context.Background(), psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{pool: pool}, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}
