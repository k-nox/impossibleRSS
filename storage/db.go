package storage

import (
	"database/sql"
	"embed"
	"fmt"
	"impossiblerss/config"
	"impossiblerss/storage/generated"

	_ "modernc.org/sqlite"

	migrate "github.com/rubenv/sql-migrate"
)

type db struct {
	conn *sql.DB
	generated.Querier
}

type DB interface {
	generated.Querier
	Migrate(embed.FS) error
}

func New(cfg *config.Database) (DB, error) {
	if cfg.Mock {
		return &mockDB{
			mockQuerier: mockQuerier{},
		}, nil
	}
	conn, err := sql.Open("sqlite", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite db: %w", err)
	}

	return &db{
		conn:    conn,
		Querier: generated.New(conn),
	}, nil
}

func (d *db) Migrate(migrationFiles embed.FS) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationFiles,
		Root:       "sqlite/migrations",
	}
	_, err := migrate.Exec(d.conn, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("error applying migrations: %w", err)
	}
	return nil
}
