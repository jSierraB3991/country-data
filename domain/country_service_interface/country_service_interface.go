package countryserviceinterface

type CountryServiceInterface interface {
	SearchCountriesAndSave() error
	HaveCountries() bool
}
