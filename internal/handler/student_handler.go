package handler

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	Services "github.com/Davethompson01/School_Paddy_golang/internal/services"
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
			studentServices,
			nil,
		)

	}
}

func LoginHandler(apiCfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var studentLogin students.Login
		if err := json.NewDecoder(r.Body).Decode(&studentLogin); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := Services.LoginInto_AsStudent(apiCfg, studentLogin)
		msg := fmt.Sprintf("Login failed: %v", err)
		if err != nil {
			RespondWithJson(w, http.StatusUnauthorized, false, msg, nil)
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
