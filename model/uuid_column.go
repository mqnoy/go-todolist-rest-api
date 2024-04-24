package model

import "github.com/google/uuid"

type UUIDColumn struct {
	ID string `gorm:"column:id;type:varchar(36);not null;primaryKey"`
}

func GenerateUUID() UUIDColumn {
	return UUIDColumn{
		ID: uuid.New().String(),
	}
}
