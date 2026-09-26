package Services

import (
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	students "github.com/Davethompson01/School_Paddy_golang/internal/models/Students"
	Studentsrepo "github.com/Davethompson01/School_Paddy_golang/internal/respositary/StudentsRepo"
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
