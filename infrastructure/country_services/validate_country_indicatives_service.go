package countryservices

import (
	"context"
	"strings"

	countryerrors "github.com/jSierraB3991/country-data/domain/country_errors"
)

func (s *CountryService) ValidateCountryIndicatives(countryId uint, countryIndicative string) error {
	indicatives, err := s.repository.FindIndicativeByCountryId(countryId)
	if err != nil {
		return err
	}
	for _, indicative := range indicatives {
		if strings.EqualFold(indicative.Indicativo, countryIndicative) {
			return nil
		}
	}
	return countryerrors.CountryIndicativeError{}
}

func (s *CountryService) ValidateCountryIndicativesCtx(ctx context.Context, countryId uint, countryIndicative string) error {
	indicatives, err := s.repository.FindIndicativeByCountryIdCtx(ctx, countryId)
	if err != nil {
		return err
	}
	for _, indicative := range indicatives {
		if strings.EqualFold(indicative.Indicativo, countryIndicative) {
			return nil
		}
	}
	return countryerrors.CountryIndicativeError{}
}
