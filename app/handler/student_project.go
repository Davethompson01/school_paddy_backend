package handler

import (
	"encoding/json"
	"net/http"
	"time"

	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	auth "github.com/Davethompson01/School_Paddy_golang/app/Auth"
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	Services "github.com/Davethompson01/School_Paddy_golang/app/services"
)

func Upload_homework(apiCfg *config.ApiConfig) http.HandlerFunc {

	return func(res http.ResponseWriter, r *http.Request) {
		var project students.Project
		if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		claims := r.Context().Value(middleware.ClaimsKey).(*auth.Claims)
		project.UserID = claims.UserID
		project.UpdatedAt = time.Now()
		uploadService, err := Services.Upload_homework(apiCfg, project)
		if err != nil {
			RespondWithJson(res, http.StatusUnauthorized, false, err.Error(), nil)
			return
		}
		RespondWithJson(res, 201, true, "Uploaded Homework", uploadService)
	}
}
