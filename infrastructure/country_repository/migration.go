package countryrepository

import (
	"context"

	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
)

func (repo *Repository) RunMigrations() error {
	if repo.db == nil {
		return nil
	}
	return repo.db.AutoMigrate(
		&countrymodels.CountryIndicatives{},
		&countrymodels.TelephoneIndicative{},
	)
}

func (repo *Repository) RunMigrationsCtx(ctx context.Context) error {
	if repo.db == nil {
		return nil
	}
	db, err := repo.GetDb(ctx)
	if err != nil {
		return err
	}
	return db.AutoMigrate(
		&countrymodels.CountryIndicatives{},
		&countrymodels.TelephoneIndicative{},
	)
}
