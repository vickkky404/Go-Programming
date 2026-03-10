package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gobackend/internal/models"
	"gobackend/internal/store"
	"gobackend/pkg/utils"
)

type UserHandler struct {
	store *store.UserStore
}

func NewUserHandler(s *store.UserStore) *UserHandler {
	return &UserHandler{store: s}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.store.GetAll()
	utils.JSONResponse(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, ok := h.store.GetByID(id)
	if !ok {
		utils.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if u.Name == "" || u.Email == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "name and email are required")
		return
	}

	created := h.store.Create(u)
	utils.JSONResponse(w, http.StatusCreated, created)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if u.Name == "" || u.Email == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "name and email are required")
		return
	}

	updated, ok := h.store.Update(id, u)
	if !ok {
		utils.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, updated)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if !h.store.Delete(id) {
		utils.ErrorResponse(w, http.StatusNotFound, "user not found")
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{"message": "user deleted"})
}
