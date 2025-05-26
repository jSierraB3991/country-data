package repositoryinterface

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

type CountryRepositoryInterface interface {
	RunMigrations() error
	SaveCountries(data []countrymodels.CountryIndicatives) error
	HaveCountries() (bool, error)
	FindCountryByCode(code string) (*countrymodels.CountryIndicatives, error)
}
