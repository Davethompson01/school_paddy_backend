package respositary

import (
	"context"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
)

func CreateStudentAccount(apiCfg *config.ApiConfig, student students.CreateStudentAccount) error {
	query := `INSERT INTO students(name, email, phone_number, password, role, auth_method) VALUES($1,$2,$3,$4,$5, $6)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int

	err := apiCfg.DB.QueryRowContext(ctx, query, student.Name, student.Email, student.Phone_Number, student.Password).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func CheckMailExist(apicfg *config.ApiConfig, email string) bool {
	var checkExist bool
	err := apicfg.DB.QueryRow(`
	SELECT EXISTS (
		SELECT 1
		FROM students
		WHERE email = $1
	)
`, email).Scan(&checkExist)
	if err != nil {
		return false
	}
	return checkExist
}

func StudentLogin(apicfg *config.ApiConfig, id int, email string, password string) (students.StudentLogin, error) {
	// var user_id int
	var user students.StudentLogin
	var userEmail string
	var userPassword string
	err := apicfg.DB.QueryRow(`
	SELECT user_id, email, password FROM students WHERE email = $1 AND password = $2
	`, email).Scan(
		&user,
		&userEmail,
		&userPassword,
	)
	if err != nil {
		return students.StudentLogin{}, nil
	}
	return user, nil
}

func GetStudentByEmail(apiCfg *config.ApiConfig, email string) (students.StudentLogin, error) {

	var user students.StudentLogin
	var userEmail string
	var userPassword string
	var userRole string
	err := apiCfg.DB.QueryRow(`
	SELECT user_id, email, password, role FROM students WHERE email = $1; `, email).Scan(
		&userEmail,
		&userPassword,
		&userRole,
	)
	if err != nil {
		return students.StudentLogin{}, err
	}
	return user, nil
}
