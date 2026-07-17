package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	Services "github.com/Davethompson01/School_Paddy_golang/app/services"
)

func StudenthandlerCreateAccount(apicfg *config.ApiConfig) http.HandlerFunc {

	return func(res http.ResponseWriter, r *http.Request) {
		var student students.CreateStudentAccount
		json.NewDecoder(r.Body).Decode(&student)

		studentServices, err := Services.CreateStudent_Service(apicfg, student)
		if err != nil {
			RespondWithJson(res, 400, false, err.Error(), nil)
			return
		}

		RespondWithJson(
			res,
			http.StatusCreated,
			true,
			"Student created successfully",
			studentServices,
		)

	}
}

func StudentLoginHandler(apiCfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var studentLogin students.StudentLogin
		if err := json.NewDecoder(r.Body).Decode(&studentLogin); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := Services.LoginInto_AsStudent(apiCfg, studentLogin)
		if err != nil {
			RespondWithJson(w, http.StatusUnauthorized, false, "Login failed", nil)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    token.AccessToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   15 * 60,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    token.RefreshToken,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			MaxAge:   7 * 24 * 60 * 60,
		})

		RespondWithJson(w, 200, true, "Login Successful", struct{}{})
	}

}
