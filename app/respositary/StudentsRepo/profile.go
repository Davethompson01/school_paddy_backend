package Studentsrepo

import (
	"github.com/Davethompson01/School_Paddy_golang/app/config"
	"github.com/Davethompson01/School_Paddy_golang/app/models"
)

func Get_Profile(apiCfg *config.ApiConfig, user_id int) error {

	var profile models.Profile
	query := `SELECT name, profile_picture, level, created_at, work_posted, role, brief_information FROM students WHERE user_id = $1`
	err := apiCfg.DB.QueryRow(query, user_id).Scan(
		&profile.Username,
		&profile.Profile_picture,
		&profile.Level,
		&profile.Created_at,
		&profile.Work_Posted,
		&profile.Role,
		&profile.Brief_infxormation,
	)

	return err

}
