package cs_prices_db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Db struct {
	Conn *pgxpool.Pool
}

func NewDatabase() (*Db, error) {
	dbUser := os.Getenv("DB_CS_USERNAME")
	dbPassword := os.Getenv("DB_CS_PASSWORD")
	dbName := os.Getenv("DB_CS_NAME")
	dbHost := os.Getenv(("DB_HOST"))
	dbPort := os.Getenv("DB_CS_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = 25
	config.MaxConnIdleTime = 5 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	return &Db{Conn: pool}, nil
}
