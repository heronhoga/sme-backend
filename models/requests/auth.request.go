package requests

type Register struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
}