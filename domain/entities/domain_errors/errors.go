package domainErrors

type InvalidField struct {
	error
	message string
}

func NewInvalidField(message string) *InvalidField {
	return &InvalidField{message: message}
}

func (err *InvalidField) Error() string {
	return err.message
}

type NotPermitted struct {
	error
	message string
}

func NewNotPermitted(message string) *NotPermitted {
	return &NotPermitted{message: message}
}

func (err *NotPermitted) Error() string {
	return err.message
}