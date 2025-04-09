package countrymodels

import "gorm.io/gorm"

type TelephoneIndicative struct {
	gorm.Model
	TelephoneIndicativeId uint   `gorm:"column:id;not null"`
	Indicativo            string `gorm:"column:indicativo;not null"`
	CountryIndicatives    CountryIndicatives
	CountryIndicativesId  uint `gorm:"column:country_id;not null"`
}
