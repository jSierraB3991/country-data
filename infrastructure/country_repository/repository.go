package countryrepository

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetConnection() *gorm.DB {
	return repo.db
}

func (repo *Repository) UpdateConnection(db *gorm.DB) {
	repo.db = db
}
