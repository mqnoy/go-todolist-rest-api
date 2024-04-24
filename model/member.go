package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	UUIDColumn
	Tasks  []Task `gorm:"many2many:MemberTask"`
	UserID string `gorm:"column:userId;type:varchar(36);"`
	User   User   `gorm:"foreignKey:UserID"`
}

func (m Member) TableName() string {
	return "Member"
}

func (m Member) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("id", uuid)

	return nil
}
