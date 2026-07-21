package Services

import (
	"fmt"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/app/models/SolutionExpert"
	"github.com/Davethompson01/School_Paddy_golang/app/respositary"
	Validation "github.com/Davethompson01/School_Paddy_golang/app/validation"
)

func Service_CreateHomeWorkBID(apiCfg *config.ApiConfig, bid solutionexpert_model.ApplyForHomeWork) (string, error) {

	err := Validation.ValidateCreateBID(bid)
	if err != nil {
		return err.Error(), err
	}
	checkIfAlreadyAccept, err := respositary.Return_Accepted_a_expert_already(apiCfg, bid.Paddyproject_id)
	if err != nil {
		return err.Error(), nil
	}
	if checkIfAlreadyAccept.Accepted_a_expert_already == true {
		return "A solution expert has already been Accepted in this project", nil
	}

	createBid := respositary.Create_Homework_BID_expert(apiCfg, bid)
	if createBid != nil {
		return createBid.Error(), nil
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
