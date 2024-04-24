package delivery

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mqnoy/go-todolist-rest-api/domain"
	"github.com/mqnoy/go-todolist-rest-api/dto"
	"github.com/mqnoy/go-todolist-rest-api/handler"
	"github.com/mqnoy/go-todolist-rest-api/pkg/clogger"
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
	request := &dto.LoginRequest{}
	if err := render.Bind(r, request); err != nil {
		handler.ParseResponse(w, r, "Login", nil, err)
		return
	}

	result, err := h.userUseCase.LoginUser(request)
	clogger.Logger().Debugf("request: %v", request)

	handler.ParseResponse(w, r, "Login", result, err)
}

func (h userHandler) Register(w http.ResponseWriter, r *http.Request) {
	handler.ParseResponse(w, r, "Register", nil, nil)
}
