package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:varchar(36);not null;index;comment:UUID" json:"id"`
	Username string `gorm:"type:varchar(64);not null;index" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

func argon2Key(password []byte, salt []byte, keyLen uint32) []byte {
	result := argon2.IDKey(password, salt, 1, 64*1024, 4, keyLen)
	return result
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
