package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/paulohfera/todo-backend-go/configs"
)

type DbContext struct {
	maxPoolSize int
	// connTimeout time.Duration
	Pool *pgxpool.Pool
}

var db *DbContext
var connection sync.Once

func NewOrGetSingleton(config *configs.Configuration) *DbContext {
	connection.Do(func() {
		conn, err := Connect(config)
		if err != nil {
			panic(err)
		}

		db = conn
	})

	return db
}

func Connect(config *configs.Configuration) (*DbContext, error) {
	db = &DbContext{
		maxPoolSize: config.Database.MaxPoolSize,
	}

	poolConfig, err := pgxpool.ParseConfig(config.URL)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
	}

	poolConfig.MaxConns = int32(db.maxPoolSize)
	db.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.NewWithConfig: %w", err)
	}

	return db, nil
}

func (p *DbContext) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
