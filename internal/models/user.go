package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	DOB   string `json:"dob" validate:"required,datetime=2006-01-02"`
}
