package countryservices

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

func (s *CountryService) FindAllCountries(orderByEnglishName bool) ([]countrymodels.CountryIndicatives, error) {
	countries, err := s.repository.FindAllCountries(orderByEnglishName)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (s *CountryService) FindIndicativeByCountryId(countryId uint) ([]countrymodels.TelephoneIndicative, error) {
	indicatives, err := s.repository.FindIndicativeByCountryId(countryId)
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}

func (s *CountryService) FindIndicativeByCounttyCode(countryCode string) ([]countrymodels.TelephoneIndicative, error) {
	indicatives, err := s.repository.FindIndicativeByCountryCode(countryCode)
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}
