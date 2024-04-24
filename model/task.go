package model

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type Task struct {
	UUIDColumn
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	DueDate     time.Time `gorm:"column:dueDate"`
	DoneAt      null.Time `gorm:"column:isDoneAt"`
	Members     []Member  `gorm:"many2many:MemberTask"`
	MemberTask  []MemberTask
	TimestampColumn
}

func (m Task) TableName() string {
	return "Task"
}

func (m Task) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewString()
	tx.Statement.SetColumn("id", uuid)

	return nil
}
