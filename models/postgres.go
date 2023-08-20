package models

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

// Open will open a sql connection with provided postgres DB
// Callers of Open need to ensure that the connection is eventually closed via db.Close() method
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())

	if err != nil {
		fmt.Errorf("open: %v", err)
	}

	return db, nil
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "3356",
		User:     "mike",
		Password: "ditka",
		Database: "galleyvalley",
		SSLMode:  "disable",
	}
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port,
		cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

// Migrate function captures the following command:
// goose postgres "host=localhost port=3356 user=mike password=ditka dbname=galleyvalley sslmode=disable" up

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")

	if err != nil {
		return fmt.Errorf("Migrate: %v", err)
	}

	err = goose.Up(db, dir)

	if err != nil {
		return fmt.Errorf("Migration Start: %v", err)
	}
	return nil
}

func MigrateFS(db *sql.DB, migrationeFs fs.FS, dir string) error {
	if dir == "" {
		dir = "."
	}
	goose.SetBaseFS(migrationeFs)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}
