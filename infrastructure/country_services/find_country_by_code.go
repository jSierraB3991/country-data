package countryservices

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

func (s *CountryService) FindCountryByIndicative(indicativeParam string) (*countrymodels.CountryIndicatives, error) {
	country, err := s.repository.FindCountryByIndicative(indicativeParam)
	if err != nil {
		return nil, err
	}
	return country, nil
}
