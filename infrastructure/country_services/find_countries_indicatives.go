package countryservices

import (
	"context"

	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
)

func (s *CountryService) FindAllCountries(orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error) {
	countries, err := s.repository.FindAllCountries(orderByEnglishName, nameOfSearchCountry)
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

func (s *CountryService) FindIndicativeByCountryCode(countryCode string) ([]countrymodels.TelephoneIndicative, error) {
	indicatives, err := s.repository.FindIndicativeByCountryCode(countryCode)
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}
func (s *CountryService) FindAllCountriesCtx(ctx context.Context, orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error) {
	countries, err := s.repository.FindAllCountriesCtx(ctx, orderByEnglishName, nameOfSearchCountry)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (s *CountryService) FindIndicativeByCountryIdCtx(ctx context.Context, countryId uint) ([]countrymodels.TelephoneIndicative, error) {
	indicatives, err := s.repository.FindIndicativeByCountryIdCtx(ctx, countryId)
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}

func (s *CountryService) FindIndicativeByCountryCodeCtx(ctx context.Context, countryCode string) ([]countrymodels.TelephoneIndicative, error) {
	indicatives, err := s.repository.FindIndicativeByCountryCodeCtx(ctx, countryCode)
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}
