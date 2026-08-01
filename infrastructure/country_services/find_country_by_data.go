package countryservices

import (
	"context"

	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
)

func (s *CountryService) FindCountryByIndicative(indicativeParam string) (*countrymodels.CountryIndicatives, error) {
	country, err := s.repository.FindCountryByIndicative(indicativeParam)
	if err != nil {
		return nil, err
	}
	return country, nil
}
func (s *CountryService) FindCountryById(countyId uint) (*countrymodels.CountryIndicatives, error) {
	country, err := s.repository.FindCountryById(countyId)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (s *CountryService) FindCountryByIndicativeCtx(ctx context.Context, indicativeParam string) (*countrymodels.CountryIndicatives, error) {
	country, err := s.repository.FindCountryByIndicativeCtx(ctx, indicativeParam)
	if err != nil {
		return nil, err
	}
	return country, nil
}
func (s *CountryService) FindCountryByIdCtx(ctx context.Context, idCountry uint) (*countrymodels.CountryIndicatives, error) {
	country, err := s.repository.FindCountryByIdCtx(ctx, idCountry)
	if err != nil {
		return nil, err
	}
	return country, nil
}
