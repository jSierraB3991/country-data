package countryrepository

import (
	"context"

	countrymodels "github.com/jSierraB3991/country-data/domain/country_models"
	eliotlibs "github.com/jSierraB3991/jsierra-libs"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *Repository) FindCountryByIndicativeCtx(ctx context.Context, indicativeParam string) (*countrymodels.CountryIndicatives, error) {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, eliotlibs.NotDatabaseConfigurateError{}
	}
	var indicative countrymodels.TelephoneIndicative
	err = db.Preload("CountryIndicatives").Where("indicativo = ?", indicativeParam).First(&indicative).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Country not found
		}
		return nil, err // Other error
	}
	return &indicative.CountryIndicatives, nil
}

func (repo *Repository) FindCountryByIdCtx(ctx context.Context, idCountry uint) (*countrymodels.CountryIndicatives, error) {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, eliotlibs.NotDatabaseConfigurateError{}
	}
	var result countrymodels.CountryIndicatives
	err = db.Where("id = ?", idCountry).First(&result).Error
	if err != nil {
		return nil, err // Other error
	}
	return &result, nil
}
func (repo *Repository) FindAllCountriesCtx(ctx context.Context, orderByEnglishName bool, nameOfSearchCountry string) ([]countrymodels.CountryIndicatives, error) {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, eliotlibs.NotDatabaseConfigurateError{}
	}
	columnOrder := "name_eng"
	if !orderByEnglishName {
		columnOrder = "name_spa"
	}
	var countries []countrymodels.CountryIndicatives
	model := db.Clauses(clause.OrderBy{Columns: []clause.OrderByColumn{{Column: clause.Column{Name: columnOrder}, Desc: false}}})

	if nameOfSearchCountry != "" {
		nameOfSearch := "name_spa"
		if orderByEnglishName {
			nameOfSearch = "name_eng"
		}
		model = model.Where("LOWER("+nameOfSearch+") LIKE LOWER(?)", "%"+nameOfSearchCountry+"%")
	}

	err = model.Find(&countries).Error
	if err != nil {
		return nil, err
	}
	return countries, nil
}
func (repo *Repository) FindIndicativeByCountryIdCtx(ctx context.Context, countryId uint) ([]countrymodels.TelephoneIndicative, error) {

	db, err := repo.GetDb(ctx)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, eliotlibs.NotDatabaseConfigurateError{}
	}
	var indicatives []countrymodels.TelephoneIndicative
	err = db.Where("country_id = ?", countryId).Find(&indicatives).Error
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}
func (repo *Repository) FindIndicativeByCountryCodeCtx(ctx context.Context, countryCode string) ([]countrymodels.TelephoneIndicative, error) {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, eliotlibs.NotDatabaseConfigurateError{}
	}
	countriesId := db.Select("id").Where("country_code = ?", countryCode).Model(&countrymodels.CountryIndicatives{})

	var indicatives []countrymodels.TelephoneIndicative
	err = db.Where("country_id IN (?)", countriesId).Find(&indicatives).Error
	if err != nil {
		return nil, err
	}
	return indicatives, nil
}
