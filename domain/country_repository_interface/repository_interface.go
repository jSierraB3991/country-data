package repositoryinterface

import (
	"context"

	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
	"gorm.io/gorm"
)

type CountryRepositoryInterface interface {
	RunMigrations() error
	RunMigrationsCtx(ctx context.Context) error

	SaveCountries(data []countrymodels.CountryIndicatives) error
	HaveCountries() (bool, error)

	FindCountryByIndicative(indicativeParam string) (*countrymodels.CountryIndicatives, error)
	FindCountryById(idCountry uint) (*countrymodels.CountryIndicatives, error)
	FindAllCountries(orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error)
	FindIndicativeByCountryId(countryId uint) ([]countrymodels.TelephoneIndicative, error)
	FindIndicativeByCountryCode(countryCode string) ([]countrymodels.TelephoneIndicative, error)

	HaveCountriesContext(ctx context.Context) (bool, error)
	SaveCountriesContext(ctx context.Context, data []countrymodels.CountryIndicatives) error
	FindCountryByIndicativeCtx(ctx context.Context, indicativeParam string) (*countrymodels.CountryIndicatives, error)
	FindCountryByIdCtx(ctx context.Context, idCountry uint) (*countrymodels.CountryIndicatives, error)
	FindAllCountriesCtx(ctx context.Context, orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error)
	FindIndicativeByCountryIdCtx(ctx context.Context, countryId uint) ([]countrymodels.TelephoneIndicative, error)
	FindIndicativeByCountryCodeCtx(ctx context.Context, countryCode string) ([]countrymodels.TelephoneIndicative, error)

	UpdateConnection(db *gorm.DB)
}
