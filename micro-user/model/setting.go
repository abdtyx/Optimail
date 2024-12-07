package model

import (
	"gorm.io/gorm"
)

type Setting struct {
	UserPK         int64  `gorm:"column:user_pk;type:bigint unsigned;not null;index" json:"user_pk"`
	Email          string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	SummaryLength  int64  `gorm:"column:summary_length;type:smallint unsigned;not null" json:"summary_length"`
	SummaryPrompt  string `gorm:"column:summary_prompt;type:text;not null" json:"summary_prompt"`
	EmphasisPrompt string `gorm:"column:emphasis_prompt;type:text;not null" json:"emphasis_prompt"`
}

func (Setting) TableName() string {
	return "user_settings"
}

func (s *Setting) BeforeCreate(_ *gorm.DB) (err error) {
	return nil
}
