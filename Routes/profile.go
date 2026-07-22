package routes

import (
	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	"github.com/Davethompson01/School_Paddy_golang/internal/handler"
	"github.com/go-chi/chi"
)

func Profile(r chi.Router, api *config.ApiConfig) {

	r.Route("/Profile", func(r chi.Router) {
		r.Use(middleware.APIKey)
		r.Use(middleware.JWTMiddleware)
		r.Use(middleware.RequireRole("admin", "super_admin", "Student", "Solution_expert"))

		r.Get("/get", handler.GetProfile_handler(api))
	})
}
