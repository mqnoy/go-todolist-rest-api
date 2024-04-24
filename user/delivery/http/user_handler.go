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

type userHandler struct {
	mux         *chi.Mux
	userUseCase domain.UserUseCase
}

func New(mux *chi.Mux, userUseCase domain.UserUseCase) {
	handler := userHandler{
		mux:         mux,
		userUseCase: userUseCase,
	}

	mux.Route("/users", func(r chi.Router) {
		r.Post("/login", handler.Login)
		r.Post("/register", handler.Register)
	})
}

func (h userHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest
	if err := render.Bind(r, &request); err != nil {
		handler.ParseResponse(w, r, "", nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	// Validate payload
	if err := cvalidator.Validator.Struct(&request); err != nil {
		handler.ParseResponse(w, r, cvalidator.ErrorValidator, nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	result, err := h.userUseCase.LoginUser(request)

	handler.ParseResponse(w, r, "Login", result, err)
}

func (h userHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request dto.RegisterRequest
	if err := render.Bind(r, &request); err != nil {
		handler.ParseResponse(w, r, "Login", nil, err)
		return
	}

	// Validate payload
	if err := cvalidator.Validator.Struct(&request); err != nil {
		handler.ParseResponse(w, r, cvalidator.ErrorValidator, nil, cerror.WrapError(http.StatusBadRequest, err))
		return
	}

	result, err := h.userUseCase.RegisterUser(request)

	handler.ParseResponse(w, r, "Register", result, err)
}
