package domain

import (
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
)

type TaskUseCase interface {
	CreateTask(param dto.CreateParam[dto.TaskCreateRequest]) (*dto.Task, error)
	UpdateTask()
	ListTasks(param dto.ListParam[dto.FilterCommonParams]) (*dto.ListResponse[dto.Task], error)
	DetailTask(param dto.DetailParam) (*dto.Task, error)
	MarkDoneTask()
	DeleteTask()
}

type TaskRepository interface {
	InsertTask(row model.Task) (*model.Task, error)
	SelectTaskById(id string) (*model.Task, error)
	SelectAndCountTask(param dto.ListParam[dto.FilterCommonParams]) (*dto.SelectAndCount[model.Task], error)
}
