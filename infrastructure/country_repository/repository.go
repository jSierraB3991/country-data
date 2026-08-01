package countryrepository

import (
	"context"
	"fmt"
	"log"

	eliotlibs "github.com/jSierraB3991/jsierra-libs"
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

func (repo *Repository) WithTenant(ctx context.Context) (*gorm.DB, error) {
	tenant, err := eliotlibs.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	tx := repo.db.Session(&gorm.Session{
		NewDB:       true,
		PrepareStmt: false,
	})

	var currentSearchPath string
	if err := tx.Raw("SHOW search_path").Scan(&currentSearchPath).Error; err != nil {
		return nil, err
	}

	if "\""+currentSearchPath+"\"" == *tenant || currentSearchPath == *tenant {
		return tx, nil
	}

	if err := tx.Exec(fmt.Sprintf(`SET search_path TO %s`, *tenant)).Error; err != nil {
		log.Printf("Error Change Schema %s", err.Error())
		return repo.WithTenant(ctx)
	}

	return tx, nil
}

func (repo *Repository) GetDb(ctx context.Context) (*gorm.DB, error) {
	db, err := repo.WithTenant(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
