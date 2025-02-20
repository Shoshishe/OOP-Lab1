package response

type UserAuthModel struct {
	FullName      string `json:"full_name" db:"full_name"`
	PasportSeries string `json:"pasport_series" db:"pasport_series"`
	PasportNum    string `json:"identif_num" db:"pasport_num"`
	MobilePhone   string `json:"phone" db:"phone"`
	Email         string `json:"email" db:"email"`
}
