package countryservices

import (
	"context"
	"log"
)

func (s *CountryService) HaveCountries() bool {
	result, err := s.repository.HaveCountries()
	if err != nil {
		log.Println(err)
		return false
	}
	return result
}

func (s *CountryService) HaveCountriesContext(ctx context.Context) bool {
	result, err := s.repository.HaveCountriesContext(ctx)
	if err != nil {
		log.Println(err)
		return false
	}
	return result
}
