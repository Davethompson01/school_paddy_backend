package Validation

import (
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStudent(student students.CreateStudentAccount) error {
	return validate.Struct(student)
}

func ValidateStudentLogin(student students.StudentLogin) error {
	return validate.Struct(student)
}
