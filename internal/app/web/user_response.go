package web

type UserResponse struct {
	ID          int64  `json:"id"`
	FullName    string `json:"fullname"`
	Email       string `json:"email"`
	NationalID  string `json:"national_id"`
	Nationality string `json:"nationality"`
}
