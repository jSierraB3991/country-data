package countryserviceinterface

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

type CountryServiceInterface interface {
	SearchCountriesAndSave() error
	HaveCountries() bool
	FindCountryByCode(code string) (*countrymodels.CountryIndicatives, error)
}
