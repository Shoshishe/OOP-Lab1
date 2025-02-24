package request

type UserAuthModel struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasportInfoModel struct {
	PasportSeries string `json:"pasport_series"`
	PasportNum    string `json:"pasport_num"`
}
type ClientSignUpModel struct {
	FullName          string `json:"full_name" db:"full_name" binding:"required"`
	PasportIdentifNum string `json:"pasport_identif_num" binding:"required"`
	MobilePhone       string `json:"phone" db:"phone_number" binding:"required"`
	PasportSeries     string `json:"pasport_series" binding:"required"`
	PasportNum        string `json:"pasport_num" binding:"required"`
	Email             string `json:"email" db:"email" binding:"required"`
	Password          string `json:"password" db:"password" binding:"required"`
}
