package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	functions "github.com/pubudu2003060/proxy_project/internal/server/functions"
	middleware "github.com/pubudu2003060/proxy_project/internal/server/middleware"
	models "github.com/pubudu2003060/proxy_project/internal/server/models"
	service "github.com/pubudu2003060/proxy_project/internal/server/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (h *UserHandler) Routes() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.UserAuthentication)

	r.Post("/", h.createUser)

	return r
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "Invalid body", err)
		return
	}

	resp, err := h.service.CreateUser(r.Context(), req)

	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	functions.RespondwithJSON(w, http.StatusCreated, resp)

}
