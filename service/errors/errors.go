package serviceErrors

type RoleError struct {
	error 
	message string
}

func (err *RoleError) Error() string {
	return err.message
}

func NewRoleError(message string) *RoleError {
	return &RoleError{message: message}
}