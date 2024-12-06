package model

import (
	"gorm.io/gorm"
)

type Emphasis struct {
	UserPK  int64  `gorm:"column:user_pk;type:bigint unsigned;not null;index" json:"user_pk"`
	Content string `gorm:"type:text;not null" json:"content"`
}

func (Emphasis) TableName() string {
	return "emphasis"
}

func (e *Emphasis) BeforeCreate(_ *gorm.DB) (err error) {
	return nil
}
