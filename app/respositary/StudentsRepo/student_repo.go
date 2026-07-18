package Studentsrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
)

func CreateStudentAccount(apiCfg *config.ApiConfig, student students.CreateStudentAccount) error {
	query := `
		INSERT INTO students(name, email, phone_number, password, role, auth_method)
		VALUES($1, $2, $3, $4, $5, $6)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := apiCfg.DB.ExecContext(
		ctx,
		query,
		student.Name,
		student.Email,
		student.Phone_Number,
		student.Password,
		student.Role,
		student.Auth_method,
	)

	return err
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

	if err == sql.ErrNoRows {
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
	// var user_id int
	// var userEmail string
	// var userPassword string
	// var userRole string
	err := apiCfg.DB.QueryRow(`
	SELECT user_id, email, password, role FROM students WHERE email = $1; `, email).Scan(
		&user.User_id,
		&user.Email,
		&user.Password,
		&user.Role,
	)
	if err != nil {
		return students.StudentLogin{}, err
	}
	return user, nil
}
