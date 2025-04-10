package mssqlApplication

import "database/sql"

type mssqlApplicationRepository struct {
	db *sql.DB
}

func NewApplicationRepository(db *sql.DB) *mssqlApplicationRepository {
	return &mssqlApplicationRepository{
		db: db,
	}
}
