package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUIDColumn
	FullName string `gorm:"column:fullName"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	TimestampColumn
}

func (m User) TableName() string {
	return "User"
}

func (m User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("id", uuid)

	return nil
}
