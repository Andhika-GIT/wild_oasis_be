package web

type CreateUser struct {
	Fullname string `validate:"required" json:"fullname"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type VerifyUser struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
