package dto

type UserRegisterDto struct {
	Name            string `json:"name" binding:"required,min=3"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=8"`
	Gender          string `json:"gender" binding:"required"`
	Dob             string `json:"dob" binding:"required,datetime=2006-01-02"`
}

type UserLoginDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
