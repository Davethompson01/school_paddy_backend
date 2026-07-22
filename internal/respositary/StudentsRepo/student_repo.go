package Studentsrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
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

	fmt.Printf("Error type: %T\n", err)
	fmt.Printf("Error value: %#v\n", err)
	return err
}

func CheckMailExist(apiCfg *config.ApiConfig, email string) bool {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1 FROM students WHERE email = $1
			UNION ALL
			SELECT 1 FROM solution_expert WHERE email = $1
		)
	`

	err := apiCfg.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func GetUserByEmail(apiCfg *config.ApiConfig, email string) (students.Login, error) {
	var user students.Login

	query := `
		SELECT user_id, email, password, role
		FROM students
		WHERE email = $1

		UNION ALL

		SELECT user_id, email, password, role
		FROM solution_expert
		WHERE email = $1

		LIMIT 1;
	`

	err := apiCfg.DB.QueryRow(query, email).Scan(
		&user.User_id,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		return students.Login{}, err
	}

	return user, nil
}
