package usecase

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/model"
	"github.com/mqnoy/go-todolist-rest-api/pkg/cerror"
	"github.com/mqnoy/go-todolist-rest-api/pkg/clogger"
	"github.com/mqnoy/go-todolist-rest-api/util"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

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

func (u *taskUseCase) CreateTask(param dto.CreateParam[dto.TaskCreateRequest]) (*dto.Task, error) {
	// TODO: Determine member from subject
	subjectId := "24a68c1b-39e9-48c7-8bf9-9ac0ad3bb312"
	member, err := u.userUseCase.GetMemberByUserId(subjectId)
	if err != nil {
		return nil, err
	}

	createValue := param.CreateValue

	dueDate, err := util.NumberToEpoch(createValue.DueDate)
	if err != nil {
		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	task := model.Task{
		Title:       createValue.Title,
		Description: createValue.Description,
		DueDate:     dueDate,
	}

	// append member for creating row on memberTask
	task.Members = append(task.Members, *member)

	taskRow, err := u.taskRepository.InsertTask(task)
	if err != nil {
		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	// Compose Response
	return u.ComposeTask(taskRow), nil
}

func (u *taskUseCase) ComposeTask(m *model.Task) *dto.Task {
	if m.ID == "" {
		return nil
	}
	return &dto.Task{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		DueDate:     util.DateToEpoch(m.DueDate),
		DoneAt:      null.Int{},
		Timestamp:   dto.ComposeTimestamp(m.TimestampColumn),
	}
}

// DeleteTask implements domain.TaskUseCase.
func (t *taskUseCase) DeleteTask() {
	panic("unimplemented")
}

func (u *taskUseCase) DetailTask(param dto.DetailParam) (*dto.Task, error) {
	// TODO: Determine member from subject
	subjectId := "24a68c1b-39e9-48c7-8bf9-9ac0ad3bb312"
	member, err := u.userUseCase.GetMemberByUserId(subjectId)
	if err != nil {
		return nil, err
	}

	taskRow, err := u.taskRepository.SelectTaskById(param.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, cerror.WrapError(http.StatusNotFound, fmt.Errorf("task not found"))
		}

		clogger.Logger().SetReportCaller(true)
		clogger.Logger().Errorf(err.Error())
		return nil, cerror.WrapError(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	}

	isOwned := u.ValidateOwnerShipTask(taskRow.MemberTask, member.ID)
	if !isOwned {
		return nil, cerror.WrapError(http.StatusForbidden, fmt.Errorf("you don't have access"))
	}

	// Compose Response
	return u.ComposeTask(taskRow), nil
}

func (u *taskUseCase) ValidateOwnerShipTask(taskMembers []model.MemberTask, memberId string) bool {
	for _, member := range taskMembers {
		if memberId == member.MemberID {
			return true
		}
	}

	return false
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
