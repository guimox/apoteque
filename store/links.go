package store

import "database/sql"

type LinksStore struct {
	db *sql.DB
}
