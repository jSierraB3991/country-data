package countryservices

import (
	repositoryinterface "github.com/jSierraB3991/country-data/domain/country_repository_interface"
	countryrepository "github.com/jSierraB3991/country-data/infrastructure/country_repository"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	"gorm.io/gorm"
)

type CountryService struct {
	urlBase    string
	repository repositoryinterface.CountryRepositoryInterface
}

func NewCountryService(database *gorm.DB, countryUrl string) *CountryService {
	repository := countryrepository.NewRepository(database)
	err := repository.RunMigrations()
	eliotlibs.FinsihApp(err)
	return &CountryService{
		repository: repository,
		urlBase:    countryUrl,
	}
}

func (s *CountryService) GetRepository() repositoryinterface.CountryRepositoryInterface {
	return s.repository
}
