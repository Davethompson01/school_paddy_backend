package solutionexpert

import (
	"context"
	"fmt"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
)

func Create_Expert_Account(apiCfg *config.ApiConfig, expert solutionexpert_model.Create_Expert_Account) error {
	query := `
		INSERT INTO solution_expert(name, email, phone_number, password, role, auth_method)
		VALUES($1, $2, $3, $4, $5, $6)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := apiCfg.DB.ExecContext(
		ctx,
		query,
		expert.Name,
		expert.Email,
		expert.Phone_Number,
		expert.Password,
		expert.Role,
		expert.Auth_method,
	)
	fmt.Printf("Name: %q, length: %d\n", expert.Name, len(expert.Name))
	fmt.Printf("Email: %q, length: %d\n", expert.Email, len(expert.Email))
	fmt.Printf("Phone: %q, length: %d\n", expert.Phone_Number, len(expert.Phone_Number))
	fmt.Printf("Password length: %d\n", len(expert.Password))
	fmt.Printf("Role: %q, length: %d\n", expert.Role, len(expert.Role))
	fmt.Printf("Auth method: %q, length: %d\n", expert.Auth_method, len(expert.Auth_method))

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
