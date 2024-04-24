package mysql

import (
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/model"
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

func (m mysqlTaskRepository) InsertTask(row model.Task) (*model.Task, error) {
	err := m.DB.Create(&row).Error
	return &row, err
}

// SelectTaskById implements domain.TaskRepository.
func (m mysqlTaskRepository) SelectTaskById(id string) (*model.Task, error) {
	var row *model.Task
	if err := m.DB.
		Preload("MemberTask").
		Where("id=?", id).First(&row).
		Error; err != nil {
		return nil, err
	}

	return row, nil
}
