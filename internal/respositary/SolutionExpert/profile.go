package solutionexpert

import (
	"context"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	"github.com/Davethompson01/School_Paddy_golang/internal/models"
)

func Get_Profile(apiCfg *config.ApiConfig, user_id int) error {

	var profile models.Profile
	query := `SELECT name, profile_picture, level, created_at, work_posted, role, brief_information FROM solution_expert WHERE user_id = $1`
	err := apiCfg.DB.QueryRow(query, user_id).Scan(
		&profile.Username,
		&profile.Profile_picture,
		&profile.Level,
		&profile.Created_at,
		&profile.Work_Posted,
		&profile.Role,
		&profile.Brief_infxormation,
	)

	return err

}

func SelectSolutionExpert_BasedOnProfile(apiCfg *config.ApiConfig, solution_expert_id int) error {
	var expert_profile models.Profile_expert
	query := `SELECT name, role, categories, level, created_at
FROM solution_expert
WHERE user_id = $1
LIMIT 10;`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := apiCfg.DB.QueryRowContext(ctx, query, solution_expert_id).Scan(
		&expert_profile.Username,
		&expert_profile.Role,
		&expert_profile.Categories,
		&expert_profile.Level,
		&expert_profile.Created_at,
	)

	return err
}
