package storage

import (
	"database/sql"
	"embed"
	"fmt"
	"impossiblerss/storage/generated"

	_ "modernc.org/sqlite"

	migrate "github.com/rubenv/sql-migrate"
)

type db struct {
	conn *sql.DB
	mock bool
	generated.Querier
}

type DB interface {
	generated.Querier
	Migrate(embed.FS) error
}

func New(dsn string, mock bool) (DB, error) {
	if mock {
		return &mockDB{
			mockQuerier: mockQuerier{},
		}, nil
	}
	conn, err := sql.Open("sqlite", dsn)
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
