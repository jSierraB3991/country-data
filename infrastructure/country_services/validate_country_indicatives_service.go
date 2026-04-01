package countryservices

import (
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
