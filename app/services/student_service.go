package Services

import (
	"errors"
	"fmt"

	auth "github.com/Davethompson01/School_Paddy_golang/app/Auth"
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	Studentsrepo "github.com/Davethompson01/School_Paddy_golang/app/respositary/StudentsRepo"
	Validation "github.com/Davethompson01/School_Paddy_golang/app/validation"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateStudent_Service(apiCfg *config.ApiConfig, studentModel students.CreateStudentAccount) (string, error) {
	exists := Studentsrepo.CheckMailExist(apiCfg, studentModel.Email)

	fmt.Println("Checking email:", studentModel.Email)
	fmt.Println("Exists:", exists)

	if exists {
		return "", errors.New("email already exists")
	}

	if err := Validation.ValidateStudent(studentModel); err != nil {
		return "", err
	}

	hashedPassword, errorHash := auth.HashPassword(studentModel.Password)
	if errorHash != nil {
		return "", errorHash
	}

	studentModel.Password = hashedPassword
	studentModel.Role = "Student"
	studentModel.Auth_method = "school_paddy_Provider"

	if err := Studentsrepo.CreateStudentAccount(apiCfg, studentModel); err != nil {

		fmt.Printf("SERVICE ERROR TYPE: %T\n", err)

		pgErr, ok := err.(*pgconn.PgError)

		fmt.Println("Type assertion:", ok)

		if ok {
			fmt.Println("Constraint:", pgErr.ConstraintName)
			fmt.Println("Code:", pgErr.Code)
		}

		return "", err
	}

	return "Student account successfully created", nil
}

func LoginInto_AsStudent(apicfg *config.ApiConfig, studentLogs students.Login) (students.LoginTokens, error) {

	checkMailExist, err := Studentsrepo.GetUserByEmail(apicfg, studentLogs.Email)
	if err != nil {
		return students.LoginTokens{}, err
	}
	// fmt.Println("%v", checkMailExist)

	err = Validation.ValidateStudentLogin(studentLogs)
	if err != nil {
		return students.LoginTokens{}, fmt.Errorf("Invalid Credentials")
	}
	comparePassword := auth.ComparePassword(checkMailExist.Password, studentLogs.Password)
	if comparePassword != nil {

		return students.LoginTokens{}, fmt.Errorf("Incorrect password")
	}

	generateToken, err := auth.GenerateToken(checkMailExist.User_id, checkMailExist.Role)
	if err != nil {
		return students.LoginTokens{}, err
	}

	refreshToken, err := auth.RefreshToken(checkMailExist.User_id, checkMailExist.Role)
	if err != nil {
		return students.LoginTokens{}, err
	}

	fmt.Println("%w", refreshToken)

	return students.LoginTokens{
		AccessToken:  generateToken,
		RefreshToken: refreshToken,
	}, nil

}
