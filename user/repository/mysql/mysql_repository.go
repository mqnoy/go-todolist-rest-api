package mysql

import (
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) domain.UserRepository {
	return &mysqlUserRepository{
		DB: db,
	}
}
