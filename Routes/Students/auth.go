package Studentsauthroute

import (
	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	"github.com/Davethompson01/School_Paddy_golang/app/handler"
	"github.com/go-chi/chi"
)

func AuthRoute(r chi.Router, apicfg *config.ApiConfig) {

	/// PUBLIC ROUTE

	r.Route("/auth", func(r chi.Router) {
		r.Use(middleware.APIKey)
		r.Post("/registerStudent", handler.StudenthandlerCreateAccount(apicfg))

		r.Post("/login", handler.StudentLoginHandler(apicfg))
	})

	//

}
