package mysql

import (
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (m mysqlTaskRepository) SelectAndCountTask(param dto.ListParam[dto.FilterCommonParams]) (*dto.SelectAndCount[model.Task], error) {
	var rows []*model.Task
	var count int64
	var result *gorm.DB

	filters := param.Filters
	orders := param.Orders
	pagination := param.Pagination
	whereClause := clause.Where{}
	mDB := m.DB

	if filters.MemberId != "" {
		mDB = mDB.Joins("JOIN MemberTask ON MemberTask.taskId = Task.id").
			Where("MemberTask.memberId = ?", filters.MemberId)
	}

	if filters.Keyword != "" {
		whereClause.Exprs = append(whereClause.Exprs, clause.Where{
			Exprs: []clause.Expression{
				clause.Like{
					Column: clause.Column{
						Name: "title",
					},
					Value: "%" + filters.Keyword + "%",
				},
			},
		})
	}

	if filters.IsDone != nil {
		if *filters.IsDone {
			whereClause.Exprs = append(whereClause.Exprs, clause.Neq{
				Column: "isDoneAt",
				Value:  nil,
			})
		} else {
			whereClause.Exprs = append(whereClause.Exprs, clause.Eq{
				Column: "isDoneAt",
				Value:  nil,
			})
		}
	}

	if len(whereClause.Exprs) > 0 {
		mDB.Clauses(whereClause)
	}

	mDB.Model(&model.Task{}).Count(&count)

	result = mDB.
		Limit(pagination.Limit).Offset(pagination.Offset).
		Order(orders).
		Find(&rows)

	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.SelectAndCount[model.Task]{
		Rows:  rows,
		Count: count,
	}, nil
}

func (m mysqlTaskRepository) UpdateTaskById(id string, values interface{}) error {
	return m.DB.Model(model.Task{}).Where("id = ?", id).Updates(values).Error
}
