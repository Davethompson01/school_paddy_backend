package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	Services "github.com/Davethompson01/School_Paddy_golang/app/services"
)

func Studenthandler(apicfg *config.ApiConfig) http.HandlerFunc {

	return func(res http.ResponseWriter, r *http.Request) {
		var student students.CreateStudentAccount
		json.NewDecoder(r.Body).Decode(&student)

		studentServices, err := Services.StudentService(apicfg, student)
		if err != nil {
			RespondWithJson(res, 400, false, err.Error(), nil)
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
