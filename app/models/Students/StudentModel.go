package students

type CreateStudentAccount struct {
    Name         string `validate:"required,min=3,max=50"`
    Email        string `validate:"required,email"`
    Phone_Number string `validate:"required,len=11,numeric"`
    Password     string `validate:"required,min=8"`
}
