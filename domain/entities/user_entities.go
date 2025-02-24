package entities

import (
	"errors"
	domainErrors "main/domain/entities/domain_errors"
)

const (
	RolePendingUser     = 0
	RoleUser            = 1
	RoleOperator        = 2
	RoleManager         = 3
	RoleOuterSpecialist = 4
	RoleAdmin           = 5
)

type UserRole = int

var rolesList []UserRole = []UserRole{
	RolePendingUser,
	RoleUser,
	RoleOperator,
	RoleManager,
	RoleManager,
	RoleOuterSpecialist,
	RoleAdmin,
}

type IdentifNum = string
type Series = string
type Password = string
type Email = string
type Phone = string
type FullName = string

type PasportInfo struct {
	pasportSeries Series
	passportNum   IdentifNum
}

func NewPasportInfo(series Series, Num IdentifNum) *PasportInfo {
	return &PasportInfo{pasportSeries: series, passportNum: Num}
}

func (info *PasportInfo) PasportSeries() Series {
	return info.pasportSeries
}

func (info *PasportInfo) PasportNum() Series {
	return info.passportNum
}

type UserOutside interface {
	CheckPassportUniqueness(pasport PasportInfo) (bool, error)
	CheckEmailUniqueness(email Email) (bool, error)
}

type User struct {
	id          int
	fullName    FullName
	pasport     PasportInfo
	mobilePhone Phone
	email       Email
	password    Password
	roleType    UserRole
	outsideInfo UserOutside
}

func NewUser(outsideInfo UserOutside, password string, email string, options ...func(*User)) (*User, error) {
	usr := &User{outsideInfo: outsideInfo}
	usr.roleType = RolePendingUser
	usr.password = password
	usr.email = email
	for _, opt := range options {
		opt(usr)
	}
	switch usr.RoleType() {
	case RolePendingUser:
		if err := usr.ValidateFullInput(); err != nil {
			return nil, err
		}
	case RoleAdmin:	
		if err := errors.Join(
			usr.ValidateEmail(),
			usr.ValidatePassword(),
			usr.ValidateFullName(),
		); err != nil {
			return nil, err
		}
	case RoleManager: 
		if err := errors.Join(
			usr.ValidateEmail(),
			usr.ValidatePassword(),
			usr.ValidateFullName(),
		); err != nil {
			return nil, err
		}
	case RoleOperator: 	
		if err := errors.Join(
			usr.ValidateEmail(),
			usr.ValidateFullName(),
			usr.ValidatePassword(),
		); err != nil {
			return nil, err
		}
	case RoleOuterSpecialist:
		if err := errors.Join(
			usr.ValidateEmail(),
			usr.ValidateFullName(),
			usr.ValidatePassword(),
			usr.ValidatePassport(),
		); err != nil {
			return nil, err
		}
	default:
		return nil, domainErrors.NewNotPermitted("incorrect role")	
	}
	return usr, nil
}

func (usr *User) ValidateFullInput() error {
	err := errors.Join(
		usr.ValidateEmail(),
		usr.ValidatePassport(),
		usr.ValidatePassword(),
		usr.ValidatePhone(),
		usr.ValidateRole(),
		usr.ValidateFullName(),
	)
	return err
}

func (usr *User) ValidateRole() error {
	if usr.roleType > len(rolesList) || usr.roleType < RolePendingUser {
		return domainErrors.NewInvalidField("invalid role input")
	}
	return nil
}

func (usr *User) ValidatePassport() error {
	if usr.pasport.pasportSeries == "" || usr.pasport.passportNum == "" {
		return domainErrors.NewInvalidField("invalid passport info")
	}
	return nil
}

func (usr *User) ValidatePassword() error {
	if len(usr.password) == 0 {
		return domainErrors.NewInvalidField("invalid password info")
	}
	return nil
}

func (usr *User) ValidateEmail() error {
	if len(usr.email) == 0 {
		return domainErrors.NewInvalidField("invalid email info")
	}
	return nil
}

func (usr *User) ValidatePhone() error {
	if len(usr.mobilePhone) == 0 {
		return domainErrors.NewInvalidField("invalid phone info")
	}
	return nil
}

func (usr *User) ValidateFullName() error {
	if len(usr.fullName) == 0 {
		return domainErrors.NewInvalidField("invalid full name")
	}
	return nil
}

func WithPhone(phone Phone) func(usr *User) {
	return func(usr *User) {
		usr.mobilePhone = phone
	}
}

func WithFullName(name FullName) func(usr *User) {
	return func(usr *User) {
		usr.fullName = name
	}
}

func WithPasportSeries(series Series) func(usr *User) {
	return func(usr *User) {
		usr.pasport.pasportSeries = series
	}
}

func WithPasportNum(pasportNum IdentifNum) func(usr *User) {
	return func(usr *User) {
		usr.pasport.passportNum = pasportNum
	}
}

func WithUserRole(usrRole UserRole) (func (usr *User)) {
	return func(usr *User) {
		usr.roleType = usrRole
	}
}

func (usr *User) Id() int {
	return usr.id
}

func (usr *User) FullName() FullName {
	return usr.fullName
}

func (usr *User) PasportSeries() Series {
	return usr.pasport.pasportSeries
}

func (usr *User) PasportNum() IdentifNum {
	return usr.pasport.passportNum
}

func (usr *User) Email() Email {
	return usr.email
}

func (usr *User) Password() Password {
	return usr.password
}

func (usr *User) MobilePhone() Phone {
	return usr.mobilePhone
}

func (usr *User) RoleType() UserRole {
	return usr.roleType
}