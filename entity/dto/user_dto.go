package dto

import "time"

type UserRegisterDto struct {
	Name            string    `json:"name" binding:"required,min=3"`
	Email           string    `json:"email" binding:"required,email"`
	Password        string    `json:"password" binding:"required,min=8"`
	ConfirmPassword string    `json:"confirm_password" binding:"required,min=8"`
	Gender          string    `json:"gender" binding:"required"`
	Dob             time.Time `json:"dob" binding:"required"`
}
