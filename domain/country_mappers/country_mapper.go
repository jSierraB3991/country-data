package countrymappers

import (
	"strings"

	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
	countryresponse "github.com/jSierraB3991/country-data/infrastructure/country_response"
)

func ToModels(countriesResponse []countryresponse.Country) []countrymodels.CountryIndicatives {
	var result []countrymodels.CountryIndicatives
	for _, r := range countriesResponse {
		if len(r.IDD.Suffixes) > 0 {
			result = append(result, toModel(r))
		}
	}
	return result
}

func toModel(country countryresponse.Country) countrymodels.CountryIndicatives {
	var indicatives []countrymodels.TelephoneIndicative
	for _, i := range country.IDD.Suffixes {
		indicatives = append(indicatives, countrymodels.TelephoneIndicative{Indicativo: country.IDD.Root + i})
	}

	lang, nativeName := getFirstNativeName(country.Name.NativeName, country.Name.Common, country.Name.Official)
	return countrymodels.CountryIndicatives{
		Name:                nativeName.Official,
		NameSpa:             country.Translations["spa"].Official,
		NameEng:             country.Name.Official,
		Language:            lang,
		FlagSvg:             country.Flags.SVG,
		TelephoneIndicative: indicatives,
		CountryCode:         strings.ToLower(country.CCA2),
	}
}

func getFirstNativeName(nativeNames map[string]countryresponse.NativeName, commonName, officialName string) (string, countryresponse.NativeName) {
	for lang, name := range nativeNames {
		return lang, name
	}
	return officialName, countryresponse.NativeName{
		Official: officialName,
		Common:   commonName,
	}
}
