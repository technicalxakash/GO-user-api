package models

type SignupRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	DOB      string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token,omitempty"`
}

type UserProfile struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	DOB   string `json:"dob"`
}

type ErrorResponse struct {
	Error struct {
		Message   string `json:"message"`
		Code      string `json:"code"`
		RequestID string `json:"request_id,omitempty"`
	} `json:"error"`
}
