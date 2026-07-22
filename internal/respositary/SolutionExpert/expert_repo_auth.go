package solutionexpert

import (
	"context"
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
		&expert.Name,
		&expert.Phone_Number,
		&expert.Password,
		&expert.Role,
		&expert.Auth_method,
	)

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
