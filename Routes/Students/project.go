package Studentsauthroute

import (
	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	"github.com/Davethompson01/School_Paddy_golang/app/handler"
	"github.com/go-chi/chi"
)

func Project(r chi.Router, apiCfg *config.ApiConfig) {

	r.Route("/project", func(r chi.Router) {
		r.Use(middleware.APIKey)
		r.Use(middleware.JWTMiddleware)
		r.Use(middleware.RequireRole("admin", "super_admin", "Student"))

		r.Post("/upload", handler.Upload_homework(apiCfg))
	})
}
