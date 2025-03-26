package models

type Role struct {
	BaseModel
	Name string `gorm:"type:string;size:30;not null;unique"`
	UserRoles *[]UserRole
}

