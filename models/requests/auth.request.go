package requests

import "github.com/google/uuid"

type Register struct {
	IdUser uuid.UUID `json:"id_user"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}