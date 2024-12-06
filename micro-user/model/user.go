package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	PK       int64  `gorm:"type:bigint unsigned;not null;auto_increment;primaryKey"`
	ID       string `gorm:"type:varchar(36);not null;index;comment:UUID" json:"id"`
	Username string `gorm:"type:varchar(64);not null;index" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (e *User) BeforeCreate(_ *gorm.DB) (err error) {
	if e.ID == "" {
		e.ID = uuid.NewString()
	}

	return nil
}
