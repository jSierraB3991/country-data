package countryresponse

type TranslationName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

type Translation struct {
	Spanish TranslationName `json:"spa"`
}

type Name struct {
	Common       string      `json:"common"`
	Official     string      `json:"official"`
	Translations Translation `json:"translations"`
}

type CountryDataResponse struct {
	Names        Name     `json:"names"`
	CallingCodes []string `json:"calling_codes"`
}
