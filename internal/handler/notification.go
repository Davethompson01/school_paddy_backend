package handler

import (
	"encoding/json"
	"net/http"

	middleware "github.com/Davethompson01/School_Paddy_golang/Middleware"
	auth "github.com/Davethompson01/School_Paddy_golang/internal/Auth"
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	Services "github.com/Davethompson01/School_Paddy_golang/internal/services"
)

func Apply_handler_notis(apiCfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var notis_ students.NotificationResponse
		if err := json.NewDecoder(r.Body).Decode(&notis_); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims := r.Context().Value(middleware.ClaimsKey).(*auth.Claims)

		notis_ForApplied, err := Services.ReturnExpertAppliedNotis(apiCfg, claims.UserID)
		if err != nil {
			RespondWithJson(w, http.StatusUnauthorized, false, err.Error(), nil)
		}
		RespondWithJson(w, 201, true, "Fetched applied notification", notis_ForApplied)

	}

}
