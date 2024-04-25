package domain

import (
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
)

type TaskUseCase interface {
	CreateTask(param dto.CreateParam[dto.TaskCreateRequest]) (*dto.Task, error)
	UpdateTask(param dto.UpdateParam[dto.TaskCreateRequest]) (*dto.Task, error)
	ListTasks(param dto.ListParam[dto.FilterCommonParams]) (*dto.ListResponse[dto.Task], error)
	DetailTask(param dto.DetailParam) (*dto.Task, error)
	MarkDoneTask(param dto.DetailParam) (*dto.Task, error)
	DeleteTask(param dto.DetailParam) error
}

type TaskRepository interface {
	InsertTask(row model.Task) (*model.Task, error)
	SelectTaskById(id string) (*model.Task, error)
	SelectAndCountTask(param dto.ListParam[dto.FilterCommonParams]) (*dto.SelectAndCount[model.Task], error)
	UpdateTaskById(id string, values interface{}) error
	DeleteTaskById(id string) error

	InsertMemberTask(row model.MemberTask) (*model.MemberTask, error)
}
