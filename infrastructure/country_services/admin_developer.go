package countryservices

import "gorm.io/gorm"

func (s *CountryService) RefreshDatabase(db *gorm.DB) {
	s.repository.UpdateConnection(db)
}
