package routes

import (
	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	"github.com/Davethompson01/School_Paddy_golang/internal/handler"
	"github.com/go-chi/chi"
)

func Project(r chi.Router, apiCfg *config.ApiConfig) {

	// STUDENTS ROUTES
	r.Route("/project", func(r chi.Router) {
		r.Use(middleware.APIKey)
		r.Use(middleware.JWTMiddleware)
		r.Use(middleware.RequireRole("admin", "super_admin", "Student"))

		//  Students
		r.Post("/upload", handler.Upload_homework(apiCfg))
		r.Post("/acceptBid", handler.HandlerAcceptBID(apiCfg))
		r.Post("/negotiateBid", handler.HandlerNegotiateBID(apiCfg))
		// r.Get("/notis_")
	})

	// SOLUTION EXPERTS ROUTES
	r.Route("/project", func(r chi.Router) {
		r.Use(middleware.APIKey)
		r.Use(middleware.JWTMiddleware)
		r.Use(middleware.RequireRole("admin", "super_admin", "Solution_expert"))

		//  SOLUTION EXPERTS
		r.Post("/createBid", handler.HandlerCreateBID(apiCfg))
		r.Post("/negotiateBid", handler.HandlerNegotiateBID(apiCfg))

	})
}
