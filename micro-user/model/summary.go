package model

import (
	"gorm.io/gorm"
)

type Summary struct {
	UserPK  int64  `gorm:"column:user_pk;type:bigint unsigned;not null;index" json:"user_pk"`
	Content string `gorm:"type:text;not null" json:"content"`
}

func (Summary) TableName() string {
	return "summary"
}

func (s *Summary) BeforeCreate(_ *gorm.DB) (err error) {
	return nil
}
