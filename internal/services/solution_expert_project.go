package Services

import (
	"fmt"

	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	rabbitmq "github.com/Davethompson01/School_Paddy_golang/internal/rabbitMQ"
	"github.com/Davethompson01/School_Paddy_golang/internal/respositary"
	Studentsrepo "github.com/Davethompson01/School_Paddy_golang/internal/respositary/StudentsRepo"
	Validation "github.com/Davethompson01/School_Paddy_golang/internal/validation"
)

func Service_CreateHomeWorkBID(apiCfg *config.ApiConfig, bid solutionexpert_model.ApplyForHomeWork) (string, error) {

	// var event solutionexpert_model.BidCreatedNotification

	err := Validation.ValidateCreateBID(bid)
	if err != nil {
		return err.Error(), err
	}
	checkIfAlreadyAccept, err := respositary.GetProjectByID(apiCfg, bid.Paddyproject_id)
	if err != nil {
		return err.Error(), nil
	}
	if checkIfAlreadyAccept.Accepted_a_expert_already == true {
		return "A solution expert has already been Accepted in this project", nil
	}
	event := solutionexpert_model.BidCreatedNotification{
		StudentID:        checkIfAlreadyAccept.Student_id,
		SolutionExpertID: bid.Solution_expert_id, // or whatever your field is called
		ProjectID:        bid.Paddyproject_id,
		Seen:             false,
	}

	createBid := respositary.Create_Homework_BID_expert(apiCfg, bid)
	if createBid != nil {
		return createBid.Error(), createBid
	}

	err = rabbitmq.PublishBidCreated(apiCfg.Rabbit.Channel, event)
	if err != nil {
		return err.Error(), err
	}

	msg := fmt.Sprintf("Solution Expert Accept Homework %v", bid.Paddyproject_id)
	return msg, nil
}

func NegotiateBid(apiCfg *config.ApiConfig, bid solutionexpert_model.NegotiateProject) (string, error) {
	err := Validation.ValidateNegotiateBID(bid)
	if err != nil {
		return err.Error(), err
	}

	negotiateBid := respositary.Negotiate_Bid(apiCfg, bid)
	if negotiateBid != nil {
		return negotiateBid.Error(), nil
	}

	return "Homework Accepted", nil
}

func StudentProjectAll(api *config.ApiConfig, studentID int) (students.ProjectSummary, error) {

	projects, err := Studentsrepo.SelectProjects(api, studentID)
	if err != nil {
		return students.ProjectSummary{}, err
	}

	summary, err := Studentsrepo.CountProjects(api, studentID)
	if err != nil {
		return students.ProjectSummary{}, err
	}

	return students.ProjectSummary{
		Projects:  projects,
		Completed: summary.Completed,
		Ongoing:   summary.Ongoing,
		Cancelled: summary.Cancelled,
	}, nil
}
