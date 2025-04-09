package countryrepository

import countrymodels "github.com/jSierraB3991/country-data/domain/country_models"

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
