package validation

import (
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	"github.com/go-playground/validator/v10"
)

type CreateStudentRequest struct {
	Name        string `validate:"required,min=3,max=50"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required,min=8"`
	PhoneNumber string `validate:"required,len=11,numeric"`
}

var validate = validator.New()

func ValidateStudent(student students.CreateStudentAccount) error {
    return validate.Struct(student)
}
