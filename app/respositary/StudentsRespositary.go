package respositary

import (
	"context"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
)

func CreateStudentAccount(apiCfg *config.ApiConfig, student students.CreateStudentAccount) error {
	query := `INSERT INTO students(name, email, phone_number, password, role) VALUES($1,$2,$3,4,$5)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int

	err := apiCfg.DB.QueryRowContext(ctx, query, student.Name, student.Email, student.Phone_Number, student.Password).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
