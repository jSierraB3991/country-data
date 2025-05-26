package countryservices

import (
	repositoryinterface "github.com/jSierraB3991/country-data/domain/country_repository_interface"
	countryrepository "github.com/jSierraB3991/country-data/infrastructure/country_repository"
	"gorm.io/gorm"
)

type CountryService struct {
	urlBase    string
	repository repositoryinterface.CountryRepositoryInterface
}

func NewCountryService(database *gorm.DB) *CountryService {
	repository := countryrepository.NewRepository(database)
	repository.RunMigrations()
	return &CountryService{
		repository: repository,
		urlBase:    "https://restcountries.com/v3.1/all",
	}
}
