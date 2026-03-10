package routes

import (
	"net/http"

	"gobackend/internal/handlers"
	"gobackend/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", middleware.Logger(http.HandlerFunc(handlers.HomeHandler)))
	mux.Handle("/health", middleware.Logger(http.HandlerFunc(handlers.HealthHandler)))
}
