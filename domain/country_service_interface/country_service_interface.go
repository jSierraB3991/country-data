package countryserviceinterface

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

type CountryServiceInterface interface {
	SearchCountriesAndSave() error
	HaveCountries() bool

	FindCountryByIndicative(indicativeParam string) (*countrymodels.CountryIndicatives, error)
	FindCountryById(countyId uint) (*countrymodels.CountryIndicatives, error)
	FindAllCountries(orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error)

	FindIndicativeByCountryId(countryId uint) ([]countrymodels.TelephoneIndicative, error)
	FindIndicativeByCounttyCode(countryCode string) ([]countrymodels.TelephoneIndicative, error)
}
