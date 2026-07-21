package Services

import (
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	"github.com/Davethompson01/School_Paddy_golang/app/respositary"
	Validation "github.com/Davethompson01/School_Paddy_golang/app/validation"
)

func Upload_homework(apiCfg *config.ApiConfig, project students.Project) (string, error) {
	err := Validation.ValidateProject(project)
	if err != nil {
		return err.Error(), err
	}
	if err := respositary.HomeWorkRespositary_IntoDB(apiCfg, project); err != nil {
		return err.Error(), err
	}
	return "Upload successful", nil
	// return
}

func AcceptBID(apiCfg *config.ApiConfig, bid students.AcceptBid_HomeWork) (string, error) {
	err := Validation.ValidateAcceptBID(bid)
	if err != nil {
		return err.Error(), err
	}

	apiCfg.DB.Begin()
	acceptBid := respositary.AcceptBid_HomeWork(apiCfg, bid)
	if acceptBid != nil {
		return acceptBid.Error(), nil
	}

	updateProjectTable := respositary.Update_paddyproject_Table_toAccept_BID(apiCfg, bid.Project_id)
	if updateProjectTable != nil {
		return updateProjectTable.Error(), nil
	}
	return "Solution expert Accept", nil
}
