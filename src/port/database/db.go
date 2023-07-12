package port

import (
	"database/sql"
	"os"

	"github.com/foliveiracamara/elixir-api-go/src/config/logger"
	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConn() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_DRIVER")
	dsn := os.Getenv("MYSQL_URL")

	db, err := sql.Open(
		dbDriver,
		dsn,
	)
	if err != nil {
		panic(err.Error())
	}

	logger.Info("Connected to database")

	return db, nil
}
