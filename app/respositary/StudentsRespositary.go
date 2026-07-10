package respositary

import (
	"context"
	"fmt"
	"time"

	"github.com/Davethompson01/School_Paddy_golang/app/config"
	students "github.com/Davethompson01/School_Paddy_golang/app/models/Students"
)

func createStudentAccount(apiCfg *config.ApiConfig, student students.CreateStudentAccount) (string, error) {
	query := `INSERT INTO students(name, email, phone_number, password) VALUES($1,$2,$3,$4)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := apiCfg.DB.QueryRowContext(ctx, query, student.Name, student.Email, student.Phone_Number, student.Password)
	if err != nil {
		return "", fmt.Errorf("Insert failed: %v", err)
	}

	return "No error Found", nil
}
