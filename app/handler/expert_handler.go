package handler

import (
	"encoding/json"

	"net/http"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/app/models/SolutionExpert"
	Services "github.com/Davethompson01/School_Paddy_golang/app/services"
)

func ExperthandlerCreateAccount(apicfg *config.ApiConfig) http.HandlerFunc {

	return func(res http.ResponseWriter, r *http.Request) {
		var expert solutionexpert_model.Create_Expert_Account
		json.NewDecoder(r.Body).Decode(&expert)

		studentServices, err := Services.CreateAccount_SolutionExpert(apicfg, expert)
		if err != nil {
			RespondWithJson(res, 400, false, err.Error(), nil)
			return
		}

		RespondWithJson(
			res,
			http.StatusCreated,
			true,
			studentServices,
			nil,
		)

	}
}

// func ExpertLoginHandler(apiCfg *config.ApiConfig) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var expertLogin solutionexpert_model.Expert_Login
// 		if err := json.NewDecoder(r.Body).Decode(&expertLogin); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		token, err := Services.LoginIn_SolutionExpert(apiCfg, expertLogin)
// 		// msg := fmt.Sprintf("Login failed: %v", err)
// 		if err != nil {
// 			RespondWithJson(w, http.StatusUnauthorized, false, "Login failed", nil)
// 			return
// 		}
// 		http.SetCookie(w, &http.Cookie{
// 			Name:     "access_token",
// 			Value:    token.AccessToken,
// 			Path:     "/",
// 			HttpOnly: true,
// 			Secure:   true,
// 			SameSite: http.SameSiteLaxMode,
// 			MaxAge:   15 * 60,
// 		})

// 		http.SetCookie(w, &http.Cookie{
// 			Name:     "refresh_token",
// 			Value:    token.RefreshToken,
// 			Path:     "/",
// 			HttpOnly: true,
// 			Secure:   true,
// 			SameSite: http.SameSiteLaxMode,
// 			MaxAge:   7 * 24 * 60 * 60,
// 		})

// 		RespondWithJson(w, 200, true, "Login Successful", struct{}{})
// 	}

// }
