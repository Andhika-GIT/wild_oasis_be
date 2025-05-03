package web

type UpdateUserNationality struct {
	NationalID  string `validate:"required" json:"national_id"`
	Nationality string `validate:"required" json:"nationality"`
}
