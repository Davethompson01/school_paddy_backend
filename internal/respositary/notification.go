package respositary

import (
	"context"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	"github.com/Davethompson01/School_Paddy_golang/internal/models"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
)

func ApplyBidNotification(apiCfg *config.ApiConfig, BidEvent solutionexpert_model.BidCreatedNotification) error {

	query := `INSERT INTO notifications
(student_id, solution_expert_id, project_id, seen, created_at, applied)
VALUES ($1, $2, $3, $4, NOW(), $5);`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := apiCfg.DB.ExecContext(
		ctx,
		query,
		BidEvent.StudentID,
		BidEvent.SolutionExpertID,
		BidEvent.ProjectID,
		BidEvent.Seen,
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
func SelectExpertApplyBidNotifications(api *config.ApiConfig, studentID int) ([]students.ExpertBidNotification, error) {
	// var
	query := `SELECT
    n.notification_id,
    n.seen,
    n.created_at,
    se.name,
    p.title
FROM notifications n
JOIN solution_expert se
    ON se.user_id = n.solution_expert_id
JOIN paddyproject p
    ON p.project_id = n.project_id
WHERE n.student_id = $1
ORDER BY n.created_at DESC
LIMIT 15;`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := api.DB.QueryContext(ctx, query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []students.ExpertBidNotification

	for rows.Next() {
		var n students.ExpertBidNotification

		err := rows.Scan(
			&n.NotificationID,
			&n.Seen,
			&n.CreatedAt,
			&n.ExpertName,
			&n.ProjectTitle,
		)
		if err != nil {
			return nil, err
		}

		notifications = append(notifications, n)
	}

	return notifications, rows.Err()

}
