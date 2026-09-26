package respositary

import (
	"context"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
)

func HomeWorkRespositary_IntoDB(apiCfg *config.ApiConfig, project students.Project) error {
	query := `
		INSERT INTO paddyproject(student_id, category, level, topic, description, bidAmount, deadline, update_at, requirement, discount_code, status)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10, $11)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := apiCfg.DB.ExecContext(ctx, query,
		project.UserID,
		project.Category,
		project.Level,
		project.Topic,
		project.Description,
		project.BidAmount,
		project.Deadline,
		project.UpdatedAt,
		project.Requirement,
		project.DiscountCode,
		project.Status)

	return err
}

func GetProjectByID(api *config.ApiConfig, project_id int) (solutionexpert_model.ApplyForHomeWork, error) {
	var project solutionexpert_model.ApplyForHomeWork
	query := `SELECT student_id, accepted_a_expert_already FROM paddyproject WHERE project_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := api.DB.QueryRowContext(ctx, query, project_id).Scan(
		&project.Student_id,
		&project.Accepted_a_expert_already,
	)
	return project, err
}

func ApprovedHomeWork(api *config.ApiConfig, project_id int) error {

	query := `UPDATE paddyproject SET completed = true WHERE project_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := api.DB.ExecContext(ctx, query, project_id)
	return err
}
