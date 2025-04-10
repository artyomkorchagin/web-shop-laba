package mssqlApplication

import "database/sql"

type ApplicationRepository struct {
	db *sql.DB
}

func NewApplicationRepository(db *sql.DB) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}
