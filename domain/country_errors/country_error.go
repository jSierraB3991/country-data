package countryerrors

type CountryIndicativeError struct{}

func (CountryIndicativeError) Error() string {
	return "COUNTRY_INDICATIVE_NOT_FOUND"
}
