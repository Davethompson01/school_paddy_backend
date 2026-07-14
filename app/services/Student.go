package services

import (
	"encoding/json"
	"net/http"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	"github.com/Davethompson01/School_Paddy_golang/app/respositary"
	"github.com/Davethompson01/School_Paddy_golang/app/validation"
)

func StudentService(apiCfg *config.ApiConfig, res http.ResponseWriter, req *http.Request) (students.CreateStudentAccount, error) {
	type Users struct {
		Name         string `json:"name"`
		Email        string `json:"email"`
		Phone_Number string `json:"phone_number"`
		Password     string `json:"password"`
	}

	var user Users

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		return students.CreateStudentAccount{}, err
	}

	student := students.CreateStudentAccount{
		Name:         user.Name,
		Email:        user.Email,
		Phone_Number: user.Phone_Number,
		Password:     user.Password,
	}
	err := validation.ValidateStudent(student)
	if err != nil {
		return students.CreateStudentAccount{}, err
	}

	insertIntoDb, err := respositary.CreateStudentAccount(apiCfg, student)
	if err != nil {
		return students.CreateStudentAccount{}, err
	}

	return insertIntoDb, nil
}
