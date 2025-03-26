package models

type Country struct {
	BaseModel
	Name string `gorm:"type:varchar(255);not null"`
	Cities []City
}


