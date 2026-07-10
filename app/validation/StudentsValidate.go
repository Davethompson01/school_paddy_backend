package validation

import (
	"github.com/go-playground/validator/v10"
)

type CreateStudentRequest struct {
	Name        string `validate:"required,min=3,max=50"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required,min=8"`
	PhoneNumber string `validate:"required,len=11,numeric"`
}

var validate = validator.New()

func ValidateStudent(req CreateStudentRequest) error {
	return validate.Struct(req)
}
