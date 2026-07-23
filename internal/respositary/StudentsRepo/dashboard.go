package Studentsrepo

import (
	"context"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	"github.com/Davethompson01/School_Paddy_golang/internal/models"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
)

func SelectProjects(api *config.ApiConfig, studentID int) ([]students.Project, error) {

	query := `
		SELECT
			project_id,
			student_id,
			category,
			level,
			topic,
			description,
			bid_amount,
			deadline,
			updated_at,
			requirement,
			status
		FROM paddyproject
		WHERE student_id = $1
		ORDER BY created_at DESC
		LIMIT 20;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := api.DB.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []students.Project

	for rows.Next() {
		var project students.Project

		err := rows.Scan(
			&project.UserID,
			&project.Category,
			&project.Level,
			&project.Topic,
			&project.Description,
			&project.BidAmount,
			&project.Deadline,
			&project.UpdatedAt,
			&project.Requirement,
			&project.Status,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
func CountProjects(api *config.ApiConfig, studentID int) (students.ProjectSummary, error) {

	var summary students.ProjectSummary

	query := `
        SELECT
            COUNT(*) FILTER (WHERE status = 'completed') AS completed,
            COUNT(*) FILTER (WHERE status = 'ongoing') AS ongoing,
            COUNT(*) FILTER (WHERE status = 'cancelled') AS cancelled
        FROM paddyproject
        WHERE student_id = $1;
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := api.DB.QueryRowContext(ctx, query, studentID).Scan(
		&summary.Completed,
		&summary.Ongoing,
		&summary.Cancelled,
	)
	if err != nil {
		return students.ProjectSummary{}, err
	}

	return summary, nil
}

func SelectExpertBasedOnCategoryRandom(api *config.ApiConfig, category string) ([]models.Profile_expert, error) {
	query := `
		SELECT
			name,
			profile_picture,
			level,
			created_at,
			work_posted,
			role,
			brief_information
		FROM solution_expert
		WHERE categories = $1
		ORDER BY RANDOM()
		LIMIT 10;
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := api.DB.QueryContext(ctx, query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experts []models.Profile_expert

	for rows.Next() {
		var expert models.Profile_expert

		err := rows.Scan(
			&expert.Username,
			&expert.Profile_picture,
			&expert.Level,
			&expert.Created_at,
			&expert.Work_Posted,
			&expert.Role,
			&expert.Brief_infxormation,
		)
		if err != nil {
			return nil, err
		}

		experts = append(experts, expert)
	}

	return experts, rows.Err()
}

// func SelectBasedRecommend() {

// }
