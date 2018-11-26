package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Service interface
type Service interface {
	GetDB() (*sql.DB, error)
}

type service struct {
	db *sql.DB
}

func (s *service) GetDB() (*sql.DB, error) {
	host := "172.17.0.1"
	port := 5432
	user := "postgres"
	password := "postgres10"
	dbname := "go-kit-test-repos"

	desc := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := ConnectToPostgres(desc)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectToPostgres func
func ConnectToPostgres(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewModelService func
func NewModelService(db *sql.DB) Service {
	return &service{
		db,
	}
}
