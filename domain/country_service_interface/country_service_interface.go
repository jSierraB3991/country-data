package countryserviceinterface

import (
	"context"

	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
	"gorm.io/gorm"
)

type CountryServiceInterface interface {
	SearchCountriesAndSave() error

	HaveCountries() bool
	FindCountryByIndicative(indicativeParam string) (*countrymodels.CountryIndicatives, error)
	FindCountryById(countryId uint) (*countrymodels.CountryIndicatives, error)
	FindAllCountries(orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error)
	FindIndicativeByCountryId(countryId uint) ([]countrymodels.TelephoneIndicative, error)
	FindIndicativeByCountryCode(countryCode string) ([]countrymodels.TelephoneIndicative, error)
	ValidateCountryIndicatives(countryId uint, countryIndicative string) error

	HaveCountriesContext(ctx context.Context) bool
	SaveCountriesContext(ctx context.Context, data []countryresponse.CountryDataResponse) error
	FindCountryByIndicativeCtx(ctx context.Context, indicativeParam string) (*countrymodels.CountryIndicatives, error)
	FindCountryByIdCtx(ctx context.Context, idCountry uint) (*countrymodels.CountryIndicatives, error)
	FindAllCountriesCtx(ctx context.Context, orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error)
	FindIndicativeByCountryIdCtx(ctx context.Context, countryId uint) ([]countrymodels.TelephoneIndicative, error)
	ValidateCountryIndicativesCtx(ctx context.Context, countryId uint, countryIndicative string) error
	FindIndicativeByCountryCodeCtx(ctx context.Context, countryCode string) ([]countrymodels.TelephoneIndicative, error)

	RefreshDatabase(db *gorm.DB)
}
