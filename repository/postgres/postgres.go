package postgres

import (
	"github.com/ericivander/monrep/config"
	"github.com/jmoiron/sqlx"

	// pq used indirectly for sqlx
	_ "github.com/lib/pq"
)

// Postgres holds database connection to postgreSQL
type Postgres struct {
	Db *sqlx.DB
}

// NewPostgres construct Postgres
func NewPostgres(cfg *config.Postgres) (*Postgres, error) {
	db, err := sqlx.Open(cfg.Driver, cfg.URL)
	if err != nil {
		return &Postgres{}, err
	}

	return &Postgres{Db: db}, nil
}
