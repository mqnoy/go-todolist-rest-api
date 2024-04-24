package usecase

import "github.com/mqnoy/go-todolist-rest-api/domain"

type taskUseCase struct {
	taskRepository domain.TaskRepository
	userUseCase    domain.UserUseCase
}

func New(taskRepository domain.TaskRepository, userUseCase domain.UserUseCase) domain.TaskUseCase {
	return &taskUseCase{
		taskRepository: taskRepository,
		userUseCase:    userUseCase,
	}
}

// CreateTask implements domain.TaskUseCase.
func (t *taskUseCase) CreateTask() {
	panic("unimplemented")
}

// DeleteTask implements domain.TaskUseCase.
func (t *taskUseCase) DeleteTask() {
	panic("unimplemented")
}

// DetailTask implements domain.TaskUseCase.
func (t *taskUseCase) DetailTask() {
	panic("unimplemented")
}

// ListTasks implements domain.TaskUseCase.
func (t *taskUseCase) ListTasks() {
	panic("unimplemented")
}

// MarkDoneTask implements domain.TaskUseCase.
func (t *taskUseCase) MarkDoneTask() {
	panic("unimplemented")
}

// UpdateTask implements domain.TaskUseCase.
func (t *taskUseCase) UpdateTask() {
	panic("unimplemented")
}
