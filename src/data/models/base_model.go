package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:timestamp with time zone;null"`
	DeletedAt sql.NullTime `gorm:"type:timestamp with time zone;null"`

	CreatedBy int `gorm:"not null"`
	ModifiedBy sql.NullInt64 `gorm:"null"`
	DeletedBy sql.NullInt64 `gorm:"null"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	userId := -1
	if value != nil {
		userId = value.(int)
	}
	base.CreatedAt = time.Now().UTC()
	base.CreatedBy = userId
	return
}

func (base *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = &sql.NullInt64{Valid: false}
	if value != nil {
		UserId = &sql.NullInt64{Int64: int64(value.(int)), Valid: true}
	}
	base.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	base.ModifiedBy = *UserId
	return
}

func (base *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = &sql.NullInt64{Valid: false}
	if value != nil {
		UserId = &sql.NullInt64{Int64: int64(value.(int)), Valid: true}
	}
	base.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	base.DeletedBy = *UserId
	return
}


