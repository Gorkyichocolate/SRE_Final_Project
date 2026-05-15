package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"sre/internal/config"
)

var DB *sql.DB

func ConnectDB() error {

	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	var err error

	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	return DB.Ping()
}
