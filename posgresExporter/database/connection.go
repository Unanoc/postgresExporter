package database

import (
	"github.com/jackc/pgx"
)

// DB is a main instance connection.
type DB struct {
	Conn *pgx.ConnPool
}

// Connect creates a connection with db.
func (db *DB) Connect(psqlURI string) error {
	pgxConfig, err := pgx.ParseURI(psqlURI)
	if err != nil {
		return err
	}

	if db.Conn, err = pgx.NewConnPool(
		pgx.ConnPoolConfig{
			ConnConfig: pgxConfig,
		}); err != nil {
		return err
	}

	return nil
}

// Disconnect closes a connection.
func (db *DB) Disconnect() {
	db.Conn.Close()
}
