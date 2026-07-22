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
		INSERT INTO paddyproject(student_id, category, level, topic, description, bidAmount, deadline, update_at, requirement, discount_code,  status)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := apiCfg.DB.ExecContext(ctx, query, project.UserID, project.Category,
		project.Level,
		project.Topic,
		project.Description,
		project.BidAmount,
		project.Deadline,
		project.UpdatedAt,
		project.Requirement,
		project.DiscountCode)

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

func Create_Homework_BID_expert(api *config.ApiConfig, apply_for_work solutionexpert_model.ApplyForHomeWork) error {
	query := `INSERT INTO applied_projects(student_id, solution_expert_id, project_id, accepted, Accepted_a_expert_already)
	VALUES($1, $2, $3, $4, $5)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := api.DB.ExecContext(ctx, query,
		apply_for_work.Student_id,
		apply_for_work.Solution_expert_id,
		apply_for_work.Paddyproject_id,
		apply_for_work.Accepted,
		apply_for_work.Accepted_a_expert_already,
	)

	return err
}

func Negotiate_Bid(api *config.ApiConfig, bid solutionexpert_model.NegotiateProject) error {
	query := `INSERT INTO negotiate(solution_expert_id, student_id, price, deadline, re_negotiate, seen)
	VALUES($1, $2, $3, $4, $5, $6)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := api.DB.ExecContext(ctx, query,
		bid.Solution_expert_id,
		bid.Student_id,
		bid.Price,
		bid.Deadline,
		bid.Renegotiate,
		bid.Seen,
	)
	return err
}

func AcceptBid_HomeWork(api *config.ApiConfig, acceptBID students.AcceptBid_HomeWork) error {
	query := `INSERT INTO accepted_bids(student_id,solution_expert_id, project_id, accepted)
	VALUES($1, $2, $3, $4)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := api.DB.ExecContext(ctx, query,
		acceptBID.Solution_expert_id,
		acceptBID.Student_id,
		acceptBID.Accepted,
	)
	return err
}

func Update_paddyproject_Table_toAccept_BID(api *config.ApiConfig, project_id int) error {
	query := `UPDATE paddyproject SET accepted_a_expert_already = true WHERE project_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := api.DB.ExecContext(ctx, query, project_id)
	return err
}

func ApprovedHomeWork(api *config.ApiConfig, project_id int) error {

	query := `UPDATE paddyproject SET completed = true WHERE project_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := api.DB.ExecContext(ctx, query, project_id)
	return err
}

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
