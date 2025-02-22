package web

type CreateUser struct {
	Username string `validate:"required" json:"username"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
