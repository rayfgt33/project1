package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type Repo struct {
	db   *sql.DB
	gorm *gorm.DB
}

type RepoInterface interface {
	BookRepo
	// BookRepo
}

// constructor function
func NewRepo(db *sql.DB, gorm *gorm.DB) *Repo {
	return &Repo{db: db, gorm: gorm}
}
