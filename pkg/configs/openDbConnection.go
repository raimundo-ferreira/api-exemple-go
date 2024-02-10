package configs

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func OpenDbConnection() (*pgx.Conn, error) {
	ctx := context.Background()

	db, err := pgx.Connect(ctx, os.Getenv("DB_SERVER_URL"))
	if err != nil {
		return nil, fmt.Errorf("Failed, not connected to database, %v", err)
	}

	return db, nil
}
