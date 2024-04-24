package domain

type TaskUseCase interface {
	CreateTask()
	UpdateTask()
	ListTasks()
	DetailTask()
	MarkDoneTask()
	DeleteTask()
}

type TaskRepository interface{}
