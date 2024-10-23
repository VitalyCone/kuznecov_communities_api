package store

import "database/sql"

type Store struct {
	config *Config
	db     *sql.DB
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}