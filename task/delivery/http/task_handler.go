package delivery

import (
	"net/http"
	"strconv"

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

func New(mux *chi.Mux, middlewareAuthorization domain.MiddlewareAuthorization, taskUseCase domain.TaskUseCase) {
	handler := taskHandler{
		mux:         mux,
		taskUseCase: taskUseCase,
	}

	mux.Route("/tasks", func(r chi.Router) {
		r.Use(middlewareAuthorization.AuthorizationJWT)
		r.Post("/", handler.PostCreateTask)
		r.Put("/{id}", handler.PutUpdateTask)
		r.Get("/", handler.GetListTasks)
		r.Get("/{id}", handler.GetDetailTask)
		r.Patch("/{id}/done", handler.PatchMarkDoneTask)
		r.Delete("/{id}", handler.DeleteTask)
	})
}

func (h taskHandler) PostCreateTask(w http.ResponseWriter, r *http.Request) {
	var request dto.TaskCreateRequest
	if err := render.Bind(r, &request); err != nil {
		handler.ParseResponse(w, r, "", err, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	// Validate payload
	if err := cvalidator.Validator.Struct(&request); err != nil {
		handler.ParseResponse(w, r, cvalidator.ErrorValidator, nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	param := dto.CreateParam[dto.TaskCreateRequest]{
		CreateValue: request,
		Session:     dto.GetAuthorizedUser(r.Context()),
	}

	// Call usecase
	result, err := h.taskUseCase.CreateTask(param)

	handler.ParseResponse(w, r, "PostCreateTask", result, err)
}

func (h taskHandler) PutUpdateTask(w http.ResponseWriter, r *http.Request) {
	var request dto.TaskCreateRequest
	if err := render.Bind(r, &request); err != nil {
		handler.ParseResponse(w, r, "", err, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	// Validate payload
	if err := cvalidator.Validator.Struct(&request); err != nil {
		handler.ParseResponse(w, r, cvalidator.ErrorValidator, nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	param := dto.UpdateParam[dto.TaskCreateRequest]{
		UpdateValue: request,
		ID:          chi.URLParam(r, "id"),
		Session:     dto.GetAuthorizedUser(r.Context()),
	}

	// Call usecase
	result, err := h.taskUseCase.UpdateTask(param)

	handler.ParseResponse(w, r, "PutUpdateTask", result, err)
}

func (h taskHandler) GetListTasks(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(handler.DefaultQuery(r, "page", "1"))
	limit, _ := strconv.Atoi(handler.DefaultQuery(r, "limit", "10"))
	offset, _ := strconv.Atoi(handler.DefaultQuery(r, "offset", "0"))
	keyword, _ := handler.GetQuery(r, "keyword")
	qIsDone, _ := handler.GetQuery(r, "isDone")
	isDone := handler.ParseQueryToBool(qIsDone)
	orders := handler.DefaultQuery(r, "orders", "id desc")

	param := dto.ListParam[dto.FilterCommonParams]{
		Filters: dto.FilterCommonParams{
			Keyword: keyword,
			IsDone:  isDone,
		},
		Orders: orders,
		Pagination: dto.Pagination{
			Page:   page,
			Limit:  limit,
			Offset: offset,
		},
		Session: dto.GetAuthorizedUser(r.Context()),
	}

	// Call usecase
	result, err := h.taskUseCase.ListTasks(param)

	handler.ParseResponse(w, r, "GetListTasks", result, err)
}
func (h taskHandler) GetDetailTask(w http.ResponseWriter, r *http.Request) {

	param := dto.DetailParam{
		ID:      chi.URLParam(r, "id"),
		Session: dto.GetAuthorizedUser(r.Context()),
	}

	// Call usecase
	result, err := h.taskUseCase.DetailTask(param)

	handler.ParseResponse(w, r, "GetDetailTask", result, err)
}

func (h taskHandler) PatchMarkDoneTask(w http.ResponseWriter, r *http.Request) {

	param := dto.DetailParam{
		ID:      chi.URLParam(r, "id"),
		Session: dto.GetAuthorizedUser(r.Context()),
	}

	// Call usecase
	result, err := h.taskUseCase.MarkDoneTask(param)

	handler.ParseResponse(w, r, "PatchMarkDoneTask", result, err)
}

func (h taskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	param := dto.DetailParam{
		ID:      chi.URLParam(r, "id"),
		Session: dto.GetAuthorizedUser(r.Context()),
	}

	// Call usecase
	err := h.taskUseCase.DeleteTask(param)

	handler.ParseResponse(w, r, "DeleteTask", nil, err)
}
