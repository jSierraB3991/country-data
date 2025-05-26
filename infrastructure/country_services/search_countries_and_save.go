package countryservices

import countryclient "github.com/jSierraB3991/country-data/infrastructure/country_client"

func (s *CountryService) SearchCountriesAndSave() error {
	data, err := countryclient.GetCountries(s.urlBase)
	if err != nil {
		return err
	}

	return s.SaveCountries(data)
}
