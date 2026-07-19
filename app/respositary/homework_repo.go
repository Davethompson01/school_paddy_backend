package respositary

import (
	"context"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
)

func HomeWorkRespositary_IntoDB(apiCfg *config.ApiConfig, project students.Project) error {
	query := `
		INSERT INTO paddyproject(user_id, category, level, topic, description, bidAmount, deadline, update_at, requirement, discount_code)
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



func ApplyForWork(){

}