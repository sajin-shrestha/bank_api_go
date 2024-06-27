package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgresql://postgres:postgres@localhost:5432/bankapi?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
