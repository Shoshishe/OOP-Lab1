package request

type UserAuthModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasportInfoModel struct {
	PasportSeries string `json:"pasport_series"`
	PasportNum    string `json:"pasport_num"`
}
type UserSignUpModel struct {
	FullName      string `json:"full_name" db:"full_name"`
	Pasport       string `json:"pasport"`
	MobilePhone   string `json:"phone" db:"phone_number"`
	PasportSeries string `json:"pasport_series"`
	PasportNum    string `json:"pasport_num"`
	Email         string `json:"email" db:"email"`
	Password      string `json:"password" db:"password"`
}
