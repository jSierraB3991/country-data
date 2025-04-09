package countrymodels

import "gorm.io/gorm"

type CountryIndicatives struct {
	gorm.Model
	CountryIndicativesId uint   `gorm:"column:id;not null"`
	Name                 string `gorm:"column:name;not null"`
	NameSpa              string `gorm:"column:name_spa;not null"`
	NameEng              string `gorm:"column:name_eng;not null"`
	FlagSvg              string `gorm:"column:flag;not null"`
	Language             string `gorm:"column:Language;not null"`
	CountryCode          string `gorm:"column:country_code;not null"`
	TelephoneIndicative  []TelephoneIndicative
}
