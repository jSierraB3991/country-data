package countryrepository

import (
	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *Repository) FindAllCountries(orderByEnglishName bool) ([]countrymodels.CountryIndicatives, error) {
	columnOrder := "name_eng"
	if !orderByEnglishName {
		columnOrder = "name_spa"
	}
	var countries []countrymodels.CountryIndicatives
	err := repo.db.Order([]clause.OrderByColumn{
		{
			Column: clause.Column{Name: columnOrder},
			Desc:   false,
		},
	}).Find(&countries).Error
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func (repo *Repository) FindIndicativeByCountryId(countryId uint) ([]countrymodels.TelephoneIndicative, error) {
	var indicatives []countrymodels.TelephoneIndicative
	err := repo.db.Where("country_id = ?", countryId).Find(&indicatives).Error
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}

func (repo *Repository) FindIndicativeByCountryCode(countryCode string) ([]countrymodels.TelephoneIndicative, error) {
	countriesId := repo.db.Select("id").Where("country_code = ?", countryCode).Model(&countrymodels.CountryIndicatives{})

	var indicatives []countrymodels.TelephoneIndicative
	err := repo.db.Where("country_id IN (?)", countriesId).Find(&indicatives).Error
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}
