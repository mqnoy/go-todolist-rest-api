package mysql

import (
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/model"
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

func (m mysqlUserRepository) SelectMemberByUserId(userId string) (*model.Member, error) {
	var row *model.Member
	if err := m.DB.
		Joins("User").Where("Member.userId=?", userId).First(&row).
		Error; err != nil {
		return nil, err
	}
	return row, nil
}

func (m mysqlUserRepository) InsertMember(row model.Member) (*model.Member, error) {
	err := m.DB.Create(&row).Error
	return &row, err
}
