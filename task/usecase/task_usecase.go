package usecase

import "github.com/mqnoy/go-todolist-rest-api/domain"

type taskUseCase struct {
	taskRepository domain.TaskRepository
}

func New(taskRepository domain.TaskRepository) domain.TaskUseCase {
	return &taskUseCase{
		taskRepository: taskRepository,
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
