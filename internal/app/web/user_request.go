package web

type UpdateUser struct {
	NationalID  string `validate:"required" json:"nationalID"`
	Nationality string `validate:"required" json:"nationality"`
}
