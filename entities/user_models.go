package entities

type User struct {
	FullName      string
	PasportSeries string
	Id            int
	MobilePhone   string
	Email         string
	Password      string
	RoleType      Role
}

type Role struct {
	name   string
	roleId int
}

func (userRole *Role) GetRoleId() int {
	return userRole.roleId
}

func (userRole *Role) GetRoleName() string {
	return userRole.name
}

var RolePendingUser = Role {
	name: "pnd_user",
	roleId: 1,
}

var RoleUser = Role{
	name: "user",
	roleId: 2,
}

var RoleOperator = Role{
	name: "operator",
	roleId: 3,
}

var RoleManager = Role{
	name: "manager",
	roleId: 4,
}

var RoleOuterSpecialis = Role{
	name: "specialist",
	roleId: 5,
}

var RoleAdmin = Role{
	name: "admin",
	roleId: 6,
}