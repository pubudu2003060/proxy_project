package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/pubudu2003060/proxy_project/internal/db/repository"
	handlers "github.com/pubudu2003060/proxy_project/internal/server/handlers"
	service "github.com/pubudu2003060/proxy_project/internal/server/service"
)

func NewRouter(db *sql.DB) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	q := repository.New(db)

	userService := service.NewUserService(q, db)
	userHandler := handlers.NewUserHandler(userService)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/user", userHandler.Routes())
	})

	return router

}
