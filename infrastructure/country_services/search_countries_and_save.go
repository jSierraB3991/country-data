package countryservices

import (
	"fmt"

	countryclient "github.com/jSierraB3991/country-data/infrastructure/country_client"
	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
)

func (s *CountryService) SearchCountriesAndSave() error {

	offset := 0
	isMore := true
	var dataResult []countryresponse.CountryDataResponse
	for isMore {
		urlFinal := fmt.Sprintf(s.urlBase, offset)
		data, err := countryclient.GetCountries(urlFinal, s.token)

		if err != nil {
			return err
		}
		offset += 100
		dataResult = append(dataResult, data.Data.Objects...)
		isMore = data.Data.Meta.More
	}

	return s.SaveCountries(dataResult)
}
