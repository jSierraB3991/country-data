package countryservices

import (
	"log"

	countrymappers "github.com/jSierraB3991/country-data/domain/country_mappers"
	repositoryinterface "github.com/jSierraB3991/country-data/domain/country_repository_interface"
	countryclient "github.com/jSierraB3991/country-data/infrastructure/country_client"
	countryrepository "github.com/jSierraB3991/country-data/infrastructure/country_repository"
	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
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

func (s *CountryService) SearchCountriesAndSave() error {
	data, err := countryclient.GetCountries(s.urlBase)
	if err != nil {
		return err
	}

	return s.SaveCountries(data)
}

func (s *CountryService) SaveCountries(data []countryresponse.Country) error {
	dataToSave := countrymappers.ToModels(data)
	return s.repository.SaveCountries(dataToSave)
}

func (s *CountryService) HaveCountries() bool {
	result, err := s.repository.HaveCountries()
	if err != nil {
		log.Println(err)
		return false
	}
	return result
}
