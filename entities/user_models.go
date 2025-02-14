package entities

const (
	RolePendingUser    = 1
	RoleUser           = 2
	RoleOperator       = 3
	RoleManager        = 4
	RoleOuterSpecialis = 5
	RoleAdmin          = 6
)

type UserRole = int
type User struct {
	FullName      string   `json:"full_name" db:"full_name"`
	PasportSeries string   `json:"pasport_series" db:"pasport_series"`
	Id            int      `json:"-" db:"id"`
	MobilePhone   string   `json:"phone" db:"phone_number"`
	Email         string   `json:"email" db:"email"`
	Password      string   `json:"password" db:"password"`
	RoleType      UserRole `json:"-" db:"role_id"`
}

func NewUser(options ...func(*User)) *User {
	usr := &User{}
	usr.RoleType = RolePendingUser
	for _, opt := range options {
		opt(usr)
	}
	return usr
}

func WithAdminRole(usr *User) {
	usr.RoleType = RoleAdmin
}

func WithManagerRole(usr *User) {
	usr.RoleType = RoleManager
}

func WithOperatorRole(usr *User) {
	usr.RoleType = RoleOperator
}

func WithOuterSpecialistRole(usr *User) {
	usr.RoleType = RoleOuterSpecialis
}
