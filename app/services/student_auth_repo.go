package Services

import (
	"errors"
	"fmt"

	auth "github.com/Davethompson01/School_Paddy_golang/app/Auth"
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	Studentsrepo "github.com/Davethompson01/School_Paddy_golang/app/respositary/StudentsRepo"
	Validation "github.com/Davethompson01/School_Paddy_golang/app/validation"
)

type LoginTokens struct {
	AccessToken  string
	RefreshToken string
}

func CreateStudent_Service(apiCfg *config.ApiConfig, studentModel students.CreateStudentAccount) (string, error) {

	if Studentsrepo.CheckMailExist(apiCfg, studentModel.Email) {
		return "", errors.New("email already exists")
	}

	if err := Validation.ValidateStudent(studentModel); err != nil {
		return "", err
	}

	hashedPassword, err := auth.HashPassword(studentModel.Password)
	if err != nil {
		return "", err
	}

	studentModel.Password = hashedPassword
	studentModel.Role = "Student"
	studentModel.Auth_method = "school_paddy_Provider"

	if err := Studentsrepo.CreateStudentAccount(apiCfg, studentModel); err != nil {
		return "", err
	}

	return "Student account successfully created", nil
}

func LoginInto_AsStudent(apicfg *config.ApiConfig, studentLogs students.StudentLogin) (LoginTokens, error) {

	checkMailExist, err := Studentsrepo.GetStudentByEmail(apicfg, studentLogs.Email)
	if err != nil {
		return LoginTokens{}, err
	}
	// fmt.Println("%v", checkMailExist)

	err = Validation.ValidateStudentLogin(studentLogs)
	if err != nil {
		return LoginTokens{}, fmt.Errorf("Invalid Credentials")
	}
	comparePassword := auth.ComparePassword(checkMailExist.Password, studentLogs.Password)
	if comparePassword != nil {

		return LoginTokens{}, fmt.Errorf("Incorrect password")
	}

	generateToken, err := auth.GenerateToken(checkMailExist.User_id, checkMailExist.Role)
	if err != nil {
		return LoginTokens{}, err
	}

	refreshToken, err := auth.RefreshToken(checkMailExist.User_id, checkMailExist.Role)
	if err != nil {
		return LoginTokens{}, err
	}

	fmt.Println("%w", refreshToken)

	return LoginTokens{
		AccessToken:  generateToken,
		RefreshToken: refreshToken,
	}, nil

}
