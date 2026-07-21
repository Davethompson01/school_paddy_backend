package handler

import (
	"net/http"

	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	auth "github.com/Davethompson01/School_Paddy_golang/app/Auth"
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	"github.com/Davethompson01/School_Paddy_golang/app/models"
	Services "github.com/Davethompson01/School_Paddy_golang/app/services"
)

func GetProfile_handler(apiCfg *config.ApiConfig) http.HandlerFunc {
	return func(res http.ResponseWriter, r *http.Request) {
		var profile models.Profile
		claims := r.Context().Value(middleware.ClaimsKey).(*auth.Claims)
		profile.User_id = claims.UserID
		profile.Role = claims.Role
		if profile.Role == "Student" {
			profile_servies_Student, err := Services.Get_Profile_student(apiCfg, claims.UserID)
			if err != nil {
				RespondWithJson(res, http.StatusUnauthorized, false, err.Error(), nil)
				return
			}
			RespondWithJson(res, 200, true, "User profile return", profile_servies_Student)
			return
		}
		if profile.Role == "Solution_expert" {
			profile_servies_expert, err := Services.Get_Profile_expert(apiCfg, claims.UserID)
			if err != nil {
				RespondWithJson(res, http.StatusUnauthorized, false, err.Error(), nil)
				return
			}
			RespondWithJson(res, 200, true, "User profile fetched", profile_servies_expert)
			return
		}

	}
}
