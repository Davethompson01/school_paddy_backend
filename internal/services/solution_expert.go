package Services

import (
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert "github.com/Davethompson01/School_Paddy_golang/internal/respositary/SolutionExpert"
)

func ExpertProfile(apiCfg *config.ApiConfig, user_id int) (string, error) {

	err := solutionexpert.Get_Profile(apiCfg, user_id)
	if err != nil {
		return err.Error(), nil
	}
	return "User_ID Profile Select", nil
}
