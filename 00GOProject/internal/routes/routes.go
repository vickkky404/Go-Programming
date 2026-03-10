package routes

import (
	"net/http"

	"gobackend/internal/handlers"
	"gobackend/internal/middleware"
	"gobackend/internal/store"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", middleware.Logger(http.HandlerFunc(handlers.HomeHandler)))
	mux.Handle("GET /health", middleware.Logger(http.HandlerFunc(handlers.HealthHandler)))

	userHandler := handlers.NewUserHandler(store.NewUserStore())

	mux.Handle("GET /users", middleware.Logger(http.HandlerFunc(userHandler.GetUsers)))
	mux.Handle("GET /users/{id}", middleware.Logger(http.HandlerFunc(userHandler.GetUser)))
	mux.Handle("POST /users", middleware.Logger(http.HandlerFunc(userHandler.CreateUser)))
	mux.Handle("PUT /users/{id}", middleware.Logger(http.HandlerFunc(userHandler.UpdateUser)))
	mux.Handle("DELETE /users/{id}", middleware.Logger(http.HandlerFunc(userHandler.DeleteUser)))
}

