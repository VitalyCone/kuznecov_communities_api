package store

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
	publicationRepository *PublicationRepository
}

func NewStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("Database is working!")

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Publication() *PublicationRepository {
	if s.publicationRepository != nil {
		return s.publicationRepository
	}

	s.publicationRepository = &PublicationRepository{
		store: s,
	}

	return s.publicationRepository
}