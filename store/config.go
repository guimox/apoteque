package store

import (
	"database/sql"
)

type Storage struct {
	Links interface {
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Links: &LinksStore{db},
	}
}
