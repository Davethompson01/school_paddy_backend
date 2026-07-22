package Validation

import (
	"errors"
	"time"

	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStudent(student students.CreateStudentAccount) error {
	if err := validate.Struct(student); err != nil {
		return FormatValidationError(err)
	}
	return nil
}

func ValidateStudentLogin(student students.Login) error {

	if err := validate.Struct(student); err != nil {
		return FormatValidationError(err)
	}
	return nil
}

func ValidateProject(project students.Project) error {

	if project.Deadline.Before(time.Now()) {
		return errors.New("deadline must be in the future")
	}

	if err := validate.Struct(project); err != nil {
		return FormatValidationError(err)
	}

	return nil
}

func ValidateExpert(expert solutionexpert_model.Create_Expert_Account) error {
	if err := validate.Struct(expert); err != nil {
		return FormatValidationError(err)
	}
	return nil
}

func ValidateExpertLogin(expert students.Login) error {

	if err := validate.Struct(expert); err != nil {
		return FormatValidationError(err)
	}
	return nil
}

func ValidateCreateBID(expert solutionexpert_model.ApplyForHomeWork) error {
	if err := validate.Struct(expert); err != nil {
		return FormatValidationError(err)
	}
	return nil
}

func ValidateNegotiateBID(expert solutionexpert_model.NegotiateProject) error {
	if err := validate.Struct(expert); err != nil {
		return FormatValidationError(err)
	}
	return nil
}

func ValidateAcceptBID(expert students.AcceptBid_HomeWork) error {
	if err := validate.Struct(expert); err != nil {
		return FormatValidationError(err)
	}
	return nil
}
