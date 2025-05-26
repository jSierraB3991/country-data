package countryservices

import "log"

func (s *CountryService) HaveCountries() bool {
	result, err := s.repository.HaveCountries()
	if err != nil {
		log.Println(err)
		return false
	}
	return result
}
