package models

type UserRole struct {
	BaseModel
	UserId uint `gorm:"not null"`
	RoleId uint `gorm:"not null"`
	User User `gorm:"foreignKey:UserId;constraint:OnDelete:NO ACTION;OnUpdate:NO ACTION"`
	Role Role `gorm:"foreignKey:RoleId;constraint:OnDelete:NO ACTION;OnUpdate:NO ACTION"`
}
