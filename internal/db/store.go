package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/EdwardKerckhof/gohtmx/config"
	"github.com/EdwardKerckhof/gohtmx/pkg/logger"
)

type Store struct {
	*Queries
	db *pgx.Conn
}

func NewStore(config config.Config, logger logger.Logger) Store {
	dbURL := formatDatabaseURL(config)
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		logger.Fatalf("failed to connect to database: %s", err.Error())
		os.Exit(1)
	}

	return Store{
		Queries: New(conn),
		db:      conn,
	}
}

func (s *Store) Close() {
	s.db.Close(context.Background())
}

func formatDatabaseURL(config config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Db.Host,
		config.Db.Port,
		config.Db.User,
		config.Db.Password,
		config.Db.Name,
	)
}
