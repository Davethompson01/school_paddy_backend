package students

type CreateStudentAccount struct {
	Name         string `validate:"required,min=3,max=50"`
	Email        string `validate:"required,email"`
	Phone_Number string `validate:"required,len=11,numeric"`
	Password     string `validate:"required,min=8"`
	Role         string `json:"role"`
	Auth_method  string `json:"auth_method"`
}

type Login struct {
	User_id  int    `json:"user_id"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
	Role     string `json:"role"`
}

type LoginTokens struct {
	AccessToken  string
	RefreshToken string
}
