package mssqlUser

import (
	"database/sql"
)

type mssqlUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *mssqlUserRepository {
	return &mssqlUserRepository{
		db: db,
	}
}
