package Services

import (
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	Studentsrepo "github.com/Davethompson01/School_Paddy_golang/internal/respositary/StudentsRepo"
)

func Get_Profile_student(apiCfg *config.ApiConfig, user_id int) (string, error) {

	err := Studentsrepo.Get_Profile(apiCfg, user_id)
	if err != nil {
		return err.Error(), nil
	}
	return "User_ID Profile Select", nil
}
