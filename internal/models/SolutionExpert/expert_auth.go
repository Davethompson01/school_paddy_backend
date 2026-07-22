package solutionexpert_model

type Create_Expert_Account struct {
	Name         string `validate:"required,min=3,max=50"`
	Email        string `validate:"required,email"`
	Phone_Number string `validate:"required,len=11,numeric"`
	Password     string `validate:"required,min=8"`
	Role         string `json:"role"`
	Auth_method  string `json:"auth_method"`
}

type LoginTokens struct {
	AccessToken  string
	RefreshToken string
}
