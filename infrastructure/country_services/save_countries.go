package countryservices

import (
	countrymappers "github.com/jSierraB3991/country-data/domain/country_mappers"
	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
)

func (s *CountryService) SaveCountries(data []countryresponse.Country) error {
	dataToSave := countrymappers.ToModels(data)
	return s.repository.SaveCountries(dataToSave)
}
