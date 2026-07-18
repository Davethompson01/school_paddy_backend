package Services

import (
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
	Studentsrepo "github.com/Davethompson01/School_Paddy_golang/app/respositary/StudentsRepo"
	Validation "github.com/Davethompson01/School_Paddy_golang/app/validation"
)

func Upload_homework(apiCfg *config.ApiConfig, project students.Project) (string, error) {
	err := Validation.ValidateProject(project)
	if err != nil {
		return "", err
	}
	if err := Studentsrepo.HomeWorkRespositary_IntoDB(apiCfg, project); err != nil {
		return "", err
	}
	return "Upload successful", nil
	// return
}
