package Services

import (
	"errors"

	auth "github.com/Davethompson01/School_Paddy_golang/app/Auth"
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	"github.com/Davethompson01/School_Paddy_golang/app/respositary"
	Validation "github.com/Davethompson01/School_Paddy_golang/app/validation"
)

type LoginTokens struct {
	AccessToken  string
	RefreshToken string
}

func CreateStudent_Service(apiCfg *config.ApiConfig, studentModel students.CreateStudentAccount) (string, error) {
	// var student students.CreateStudentAccount

	checkMailExist := respositary.CheckMailExist(apiCfg, studentModel.Email)
	if checkMailExist == true {
		return "Email already exist", nil
	}
	err := Validation.ValidateStudent(studentModel)
	if err != nil {
		return err.Error(), err
	}

	hashedPassword, err := auth.HashPassword(studentModel.Password)
	if err != nil {
		return err.Error(), err
	}
	studentModel.Password = hashedPassword
	studentModel.Role = "Student"
	studentModel.Auth_method = "school_paddy_Provider"
	errRespositary := respositary.CreateStudentAccount(apiCfg, studentModel)
	if errRespositary != nil {
		return "Failed to create user", errRespositary
	}

	return "Student Account Successfuly created", nil
}

func LoginInto_AsStudent(apicfg *config.ApiConfig, studentLogs students.StudentLogin) (LoginTokens, error) {

	checkMailExist, err := respositary.GetStudentByEmail(apicfg, studentLogs.Email)
	if err != nil {
		return LoginTokens{}, errors.New("email doesn't exist")
	}
	err = Validation.ValidateStudentLogin(studentLogs)
	if err != nil {
		return LoginTokens{}, err
	}
	comparePassword := auth.ComparePassword(checkMailExist.Password, studentLogs.Password)
	if comparePassword != nil {
		return LoginTokens{}, errors.New("invalid password")
	}

	generateToken, err := auth.GenerateToken(checkMailExist.User_id, checkMailExist.Role)
	if err != nil {
		return LoginTokens{}, err
	}

	refreshToken, err := auth.RefreshToken(checkMailExist.User_id, checkMailExist.Role)
	if err != nil {
		return LoginTokens{}, err
	}

	return LoginTokens{
		AccessToken:  generateToken,
		RefreshToken: refreshToken,
	}, nil

}
