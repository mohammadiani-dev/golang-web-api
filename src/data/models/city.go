package models

type City struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null"`
	CountryId int `gorm:"not null"`
	Country Country `gorm:"foreignKey:CountryId"`
}


