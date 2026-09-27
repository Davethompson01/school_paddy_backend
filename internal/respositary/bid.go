package respositary

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
)

func CreateBid(
	api *config.ApiConfig,
	apply_for_work solutionexpert_model.ApplyForHomeWork,
) (int, error) {

	query := `
        INSERT INTO bid(
            student_id,
            solution_expert_id,
            project_id,
            accepted,
            accepted_a_expert_already,
            isCompleted,
            status
        )
        VALUES($1, $2, $3, $4, $5, $6, $7)
        RETURNING bid_id
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var bidID int

	err := api.DB.QueryRowContext(
		ctx,
		query,
		apply_for_work.Student_id,
		apply_for_work.Solution_expert_id,
		apply_for_work.Paddyproject_id,
		apply_for_work.Accepted,
		apply_for_work.Accepted_a_expert_already,
		apply_for_work.IsCompleted,
		apply_for_work.Status,
	).Scan(&bidID)

	if err != nil {
		return 0, err
	}

	return bidID, nil
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

func AcceptBid_HomeWork(api *config.ApiConfig, acceptBID students.AcceptBid) error {
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
	// x := "string"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := api.DB.ExecContext(ctx, query, project_id)
	return err
}

// func 
func BidDetails(
	api *config.ApiConfig,
	userID, projectID, student_id int,
) ([]solutionexpert_model.ProjectDetails, error) {

	query := `
    SELECT
        details.name,
			COALESCE(details.profile_pics, '') AS profile_pics,
        COALESCE(details.categories, '') AS categories,
        details.created_at,
        COALESCE(details.reviews, 0) AS reviews,
        project.deadline,
        project.bidamount
    FROM solution_expert AS details
    CROSS JOIN paddyproject AS project
    WHERE details.user_id = $1
      AND project.project_id = $2
      AND project.student_id = $3;
`
	ctx, cancel := context.WithTimeout(
		context.Background(),
		3*time.Second,
	)
	defer cancel()
	fmt.Print(userID, projectID, student_id)

	rows, err := api.DB.QueryContext(
		ctx,
		query,
		userID,
		projectID,
		student_id,
	)

	if err != nil {
		return []solutionexpert_model.ProjectDetails{}, err
	}

	defer rows.Close()

	detail := make([]solutionexpert_model.ProjectDetails, 0)

	for rows.Next() {

		var details solutionexpert_model.ProjectDetails

		err := rows.Scan(
			&details.Username,
			&details.Profile_Pics,
			&details.Categories,
			&details.CreatedAt,
			&details.Review,
			&details.Deadline,
			&details.Price,
		)
		log.Print(err)

		if err != nil {
			return []solutionexpert_model.ProjectDetails{}, fmt.Errorf(
				"failed to scan project details: %w",
				err,
			)
		}

		detail = append(detail, details)
	}

	if err := rows.Err(); err != nil {
		return []solutionexpert_model.ProjectDetails{}, fmt.Errorf(
			"failed while reading project details: %w",
			err,
		)
	}

	log.Print(detail)
	return detail, nil
}

func CheckBidExistsAndOwner(api *config.ApiConfig, bidID, userID int) bool {

	var exists bool
	query := `SELECT EXISTS(
		SELECT 1 FROM bid where bid_id = $1
		UNION ALL
		SELECT 1 FROM bid where student_id = $2
	)`

	err := api.DB.QueryRow(query, bidID, userID).Scan(&exists)
	if err != nil {
		return false
	}
	// fmt.Sprint(exists)

	return exists
}



// func CheckHomeworkStatus(api *config.ApiConfig, project){
//     query := `SELECT student_id, category, status from paddyproject WHERE project_id = $1`

// }