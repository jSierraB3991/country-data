package countryrepository

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

func (repo *Repository) RunMigrations() error {
	return repo.db.AutoMigrate(
		&countrymodels.CountryIndicatives{},
		&countrymodels.TelephoneIndicative{},
	)
}
