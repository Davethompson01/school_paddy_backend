package Services

import (
	"errors"
	"fmt"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	rabbitmq "github.com/Davethompson01/School_Paddy_golang/internal/rabbitMQ"
	"github.com/Davethompson01/School_Paddy_golang/internal/respositary"
	Validation "github.com/Davethompson01/School_Paddy_golang/internal/validation"
)

func BidDetails(api *config.ApiConfig, userId, projectID, bidID int, solutionExpert int) ([]solutionexpert_model.ProjectDetails, error) {

	// check if Bids exists and project owner
	bidExists := respositary.CheckBidExistsAndOwner(api, bidID, userId)
	if !bidExists {
		return []solutionexpert_model.ProjectDetails{}, fmt.Errorf("Failed to Unauthorize Current user")
	}

	// owner of the project

	// you can't create multiple bid with same users
	projectDetail, err := respositary.BidDetails(api, solutionExpert, projectID, userId)
	// fmt.Printf("Project ID %v", projectID)
	if err != nil {
		return []solutionexpert_model.ProjectDetails{}, fmt.Errorf("Failed to load details %v", err)
	}
	// log.Print(projectDetail)
	return projectDetail, nil
}

func CreateBid(
	apiCfg *config.ApiConfig,
	bid solutionexpert_model.ApplyForHomeWork,
) (string, error) {

	err := Validation.ValidateCreateBID(bid)
	if err != nil {
		return err.Error(), err
	}

	project, err := respositary.GetProjectByID(
		apiCfg,
		bid.Paddyproject_id,
	)
	if err != nil {
		return err.Error(), err
	}

	if project.Accepted_a_expert_already {
		return "A solution expert has already been Accepted in this project", nil
	}

	bid.Accepted = false
	bid.IsCompleted = false
	bid.Status = "Pending"

	// Create BID in PostgreSQL
	BidID, err := respositary.CreateBid(apiCfg, bid)
	if err != nil {
		return err.Error(), err
	}

	if apiCfg.Rabbit == nil {
		return "rabbitmq is not initialized",
			errors.New("rabbitmq is not initialized")
	}

	if apiCfg.Rabbit.Channel == nil {
		return "rabbitmq channel is not initialized",
			errors.New("rabbitmq channel is not initialized")
	}

	// Create RabbitMQ event
	event := solutionexpert_model.BidCreatedNotification{
		StudentID:        project.Student_id,
		SolutionExpertID: bid.Solution_expert_id,
		ProjectID:        bid.Paddyproject_id,
		BidID:            BidID,
		Message:          fmt.Sprintf("Solution Expert Created a Bid, %v", bid.Status),
	}

	// Put event into RabbitMQ
	err = rabbitmq.PublishBidCreated(
		apiCfg.Rabbit.Channel,
		event,
	)

	if err != nil {
		return err.Error(), err
	}

	msg := fmt.Sprintf(
		"Solution Expert Created a Bid %v",
		bid.Status,
	)

	return msg, nil
}

// func NegotiateBid(apiCfg *config.ApiConfig, bid solutionexpert_model.NegotiateProject) (string, error) {
// 	err := Validation.ValidateNegotiateBID(bid)
// 	if err != nil {
// 		return err.Error(), err
// 	}

// 	negotiateBid := respositary.Negotiate_Bid(apiCfg, bid)
// 	if negotiateBid != nil {
// 		return negotiateBid.Error(), nil
// 	}

// 	return "Homework Accepted", nil
// }

func AcceptBID(apiCfg *config.ApiConfig, bid students.AcceptBid) (string, error) {
	err := Validation.ValidateAcceptBID(bid)
	if err != nil {
		return err.Error(), err
	}

	checkBidStatus, err := respositary.GetProjectByID(apiCfg, bid.Project_id)
	if err != nil {
		return checkBidStatus.Status, err
	}

	if checkBidStatus.Accepted_a_expert_already {
		return "Negotiation ongoing on project already", err
	}

	apiCfg.DB.Begin()
	acceptBid := respositary.AcceptBid_HomeWork(apiCfg, bid)
	if acceptBid != nil {
		return acceptBid.Error(), acceptBid
	}

	updateProjectTable := respositary.Update_paddyproject_Table_toAccept_BID(apiCfg, bid.Project_id)
	if updateProjectTable != nil {
		return updateProjectTable.Error(), updateProjectTable
	}




	return "Solution expert Accept", nil
}
