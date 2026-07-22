package Services

import (
	"errors"

	auth "github.com/Davethompson01/School_Paddy_golang/internal/Auth"
	"github.com/Davethompson01/School_Paddy_golang/internal/config"
	solutionexpert_model "github.com/Davethompson01/School_Paddy_golang/internal/models/SolutionExpert"
	solutionexpert "github.com/Davethompson01/School_Paddy_golang/internal/respositary/SolutionExpert"
	Validation "github.com/Davethompson01/School_Paddy_golang/internal/validation"
)

func CreateAccount_SolutionExpert(apiCfg *config.ApiConfig, expertModel solutionexpert_model.Create_Expert_Account) (string, error) {

	if solutionexpert.CheckMailExist(apiCfg, expertModel.Email) {
		return "", errors.New("email already exists")
	}

	if err := Validation.ValidateExpert(expertModel); err != nil {
		return "", err
	}

	hashedPassword, err := auth.HashPassword(expertModel.Password)
	if err != nil {
		return "", err
	}

	expertModel.Password = hashedPassword
	expertModel.Role = "Solution_Expert"
	expertModel.Auth_method = "school_paddy_Provider"

	if err := solutionexpert.Create_Expert_Account(apiCfg, expertModel); err != nil {
		return "", err
	}

	return "Solution Expert successfully created", nil
}

// func LoginIn_SolutionExpert(apicfg *config.ApiConfig, expertLogs students.Login) (solutionexpert_model.LoginTokens, error) {

// 	checkMailExist, err := solutionexpert.GetUserByEmail(apicfg, expertLogs.Email)
// 	if err != nil {
// 		return solutionexpert_model.LoginTokens{}, err
// 	}
// 	// fmt.Println("%v", checkMailExist)
// 	err = Validation.ValidateExpertLogin(expertLogs)
// 	if err != nil {
// 		return solutionexpert_model.LoginTokens{}, fmt.Errorf("Invalid Credentials")
// 	}
// 	comparePassword := auth.ComparePassword(checkMailExist.Password, expertLogs.Password)
// 	if comparePassword != nil {

// 		return solutionexpert_model.LoginTokens{}, fmt.Errorf("Incorrect password")
// 	}

// 	generateToken, err := auth.GenerateToken(checkMailExist.User_id, checkMailExist.Role)
// 	if err != nil {
// 		return solutionexpert_model.LoginTokens{}, err
// 	}

// 	refreshToken, err := auth.RefreshToken(checkMailExist.User_id, checkMailExist.Role)
// 	if err != nil {
// 		return solutionexpert_model.LoginTokens{}, err
// 	}

// 	return solutionexpert_model.LoginTokens{
// 		AccessToken:  generateToken,
// 		RefreshToken: refreshToken,
// 	}, nil

// }
