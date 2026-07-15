package Services

import (
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	"github.com/Davethompson01/School_Paddy_golang/app/respositary"
	Validation "github.com/Davethompson01/School_Paddy_golang/app/validation"
)

func StudentService(apiCfg *config.ApiConfig, studentModel students.CreateStudentAccount) (students.CreateStudentAccount, error) {
	// var student students.CreateStudentAccount

	err := Validation.ValidateStudent(studentModel)
	if err != nil {
		return studentModel, err
	}
	errRespositary := respositary.CreateStudentAccount(apiCfg, studentModel)
	if errRespositary != nil {

		return students.CreateStudentAccount{}, err
	}

	return studentModel, nil
}
