package app

import (
	"example-go/controller"

	"github.com/joinpickup/middleware-go/logging"

	"github.com/go-chi/chi/v5"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	// middleware
	r.Use(logging.LoggerMiddleware(&logging.Logger))
	r.Get("/", controller.GetHealth)

	// setup routes
	r.Route("/v1", func(r chi.Router) {
		// health routes
		r.Get("/health", controller.GetHealth)
		r.Get("/healthping", controller.GetPing)
	})

	return r
}
