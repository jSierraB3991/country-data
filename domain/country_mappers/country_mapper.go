package countrymappers

import (
	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
)

func ToModels(countriesResponse []countryresponse.CountryDataResponse) []countrymodels.CountryIndicatives {
	var result []countrymodels.CountryIndicatives
	for _, r := range countriesResponse {
		result = append(result, countrymodels.CountryIndicatives{
			Name:                r.Names.Common,
			NameSpa:             r.Names.Translations.Spanish.Common,
			NameEng:             r.Names.Official,
			FlagSvg:             "",                             // Assuming this field is not available in the response
			TelephoneIndicative: getIndicatives(r.CallingCodes), // Assuming this field is not available in the response
			CountryCode:         "",                             // Assuming this field is not available in the response
		})
	}
	return result
}

func getIndicatives(callingCodes []string) []countrymodels.TelephoneIndicative {
	var indicatives []countrymodels.TelephoneIndicative
	for _, code := range callingCodes {
		indicatives = append(indicatives, countrymodels.TelephoneIndicative{
			Indicativo: code,
		})
	}
	return indicatives
}
