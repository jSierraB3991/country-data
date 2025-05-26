package countryrepository

import (
	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
	"gorm.io/gorm"
)

func (repo *Repository) SaveCountries(data []countrymodels.CountryIndicatives) error {
	return repo.db.Model(&countrymodels.CountryIndicatives{}).Save(&data).Error
}

func (repo *Repository) HaveCountries() (bool, error) {
	var countriesCount int64
	err := repo.db.Model(&countrymodels.CountryIndicatives{}).Count(&countriesCount).Error
	if err != nil {
		return false, err
	}
	return countriesCount > 0, nil
}

func (repo *Repository) FindCountryByIndicative(indicativeParam string) (*countrymodels.CountryIndicatives, error) {
	var indicative countrymodels.TelephoneIndicative
	err := repo.db.Preload("CountryIndicatives").Where("indicativo = ?", indicativeParam).First(&indicative).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Country not found
		}
		return nil, err // Other error
	}
	return &indicative.CountryIndicatives, nil
}
