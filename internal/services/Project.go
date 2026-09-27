package Services

import (
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	"github.com/Davethompson01/School_Paddy_golang/internal/respositary"
	Studentsrepo "github.com/Davethompson01/School_Paddy_golang/internal/respositary/StudentsRepo"
	Validation "github.com/Davethompson01/School_Paddy_golang/internal/validation"
)

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

func Upload_homework(apiCfg *config.ApiConfig, project students.Project) (string, error) {
	err := Validation.ValidateProject(project)
	if err != nil {
		return err.Error(), err
	}

	project.Status = "Pending"
	if err := respositary.HomeWorkRespositary_IntoDB(apiCfg, project); err != nil {
		return err.Error(), err
	}
	return "Upload successful", nil
	// return
}
