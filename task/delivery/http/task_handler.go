package delivery

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/handler"
	"github.com/mqnoy/go-todolist-rest-api/pkg/cerror"
	"github.com/mqnoy/go-todolist-rest-api/pkg/cvalidator"
)

type taskHandler struct {
	mux         *chi.Mux
	taskUseCase domain.TaskUseCase
}

func New(mux *chi.Mux, taskUseCase domain.TaskUseCase) {
	handler := taskHandler{
		mux:         mux,
		taskUseCase: taskUseCase,
	}

	mux.Route("/tasks", func(r chi.Router) {
		r.Post("/", handler.PostCreateTask)
		r.Put("/:id", handler.PutUpdateTask)
		r.Get("/", handler.GetListTasks)
		r.Get("/{id}", handler.GetDetailTask)
		r.Patch("/:id/done", handler.PatchMarkDoneTask)
		r.Delete("/:id", handler.DeleteTask)
	})
}

func (h taskHandler) PostCreateTask(w http.ResponseWriter, r *http.Request) {
	request := &dto.TaskCreateRequest{}
	if err := render.Bind(r, request); err != nil {
		handler.ParseResponse(w, r, "", err, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	// Validate payload
	if err := cvalidator.Validator.Struct(request); err != nil {
		handler.ParseResponse(w, r, cvalidator.ErrorValidator, nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	param := dto.CreateParam[dto.TaskCreateRequest]{
		CreateValue: *request,
	}

	// Call usecase
	result, err := h.taskUseCase.CreateTask(param)

	handler.ParseResponse(w, r, "PostCreateTask", result, err)
}

func (h taskHandler) PutUpdateTask(w http.ResponseWriter, r *http.Request) {
	handler.ParseResponse(w, r, "PutUpdateTask", nil, nil)
}

func (h taskHandler) GetListTasks(w http.ResponseWriter, r *http.Request) {
	handler.ParseResponse(w, r, "GetListTasks", nil, nil)
}

func (h taskHandler) GetDetailTask(w http.ResponseWriter, r *http.Request) {

	param := dto.DetailParam{
		ID: chi.URLParam(r, "id"),
	}

	// Call usecase
	result, err := h.taskUseCase.DetailTask(param)

	handler.ParseResponse(w, r, "GetDetailTask", result, err)
}

func (h taskHandler) PatchMarkDoneTask(w http.ResponseWriter, r *http.Request) {
	handler.ParseResponse(w, r, "PatchMarkDoneTask", nil, nil)
}

func (h taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	handler.ParseResponse(w, r, "DeleteTask", nil, nil)
}
