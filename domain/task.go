package domain

import (
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
)

type TaskUseCase interface {
	CreateTask(param dto.CreateParam[dto.TaskCreateRequest]) (*dto.Task, error)
	UpdateTask()
	ListTasks()
	DetailTask()
	MarkDoneTask()
	DeleteTask()
}

type TaskRepository interface {
	InsertTask(row model.Task) (*model.Task, error)
}
