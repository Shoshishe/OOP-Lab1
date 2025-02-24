package usersServices

type AdminService interface {
	CancelOperation(operationId, userId int) error
}