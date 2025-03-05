package storage

import (
	"database/sql"
	"embed"
	"fmt"
	"impossiblerss/storage/generated"

	_ "modernc.org/sqlite"

	migrate "github.com/rubenv/sql-migrate"
)

type DB struct {
	conn *sql.DB
	generated.Querier
}

func New(dsn string) (*DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening sqlite db: %w", err)
	}

	return &DB{
		conn:    db,
		Querier: generated.New(db),
	}, nil
}

func (db *DB) Migrate(migrationFiles embed.FS) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationFiles,
		Root:       "sqlite/migrations",
	}
	_, err := migrate.Exec(db.conn, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("error applying migrations: %w", err)
	}
	return nil
}
