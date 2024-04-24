package mysql

import (
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"gorm.io/gorm"
)

type mysqlTaskRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) domain.TaskRepository {
	return &mysqlTaskRepository{
		DB: db,
	}
}
