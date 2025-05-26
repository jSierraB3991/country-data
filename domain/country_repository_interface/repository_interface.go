package repositoryinterface

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

type CountryRepositoryInterface interface {
	RunMigrations() error
	SaveCountries(data []countrymodels.CountryIndicatives) error
	HaveCountries() (bool, error)
	FindCountryByIndicative(indicative string) (*countrymodels.CountryIndicatives, error)
	FindAllCountries() ([]countrymodels.CountryIndicatives, error)
	FindIndicativeByCountryId(countryId uint) ([]countrymodels.TelephoneIndicative, error)
}
