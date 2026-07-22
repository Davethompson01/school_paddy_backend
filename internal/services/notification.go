package Services

import (
	"fmt"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	"github.com/Davethompson01/School_Paddy_golang/internal/respositary"
)

func ReturnExpertAppliedNotis(api *config.ApiConfig, studentID int) ([]students.NotificationResponse, error) {

	experts, err := respositary.SelectExpertApplyBidNotifications(api, studentID)
	if err != nil {
		return nil, err
	}

	var responses []students.NotificationResponse

	for _, expert := range experts {

		message := fmt.Sprintf(
			"%s applied to your homework project.",
			expert.ExpertName,
		)

		responses = append(responses, students.NotificationResponse{
			ID:        expert.NotificationID,
			Message:   message,
			Seen:      expert.Seen,
			CreatedAt: expert.CreatedAt,
		})
	}

	return responses, nil
}
