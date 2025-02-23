package usersServices

import (
	"main/service"
	"main/service/entities_models/request"
)

type OuterSpecialistService interface {
	SendInfoForPayment(request.PaymentRequestModel) error
	TransferRequest(transfer request.TransferModel) error
}

type OuterSpecialistServiceImpl struct {
	OuterSpecialistService
	repos service.OuterSpecialistRepository
}

func NewOuterSpecialistService(repos service.OuterSpecialistRepository) *OuterSpecialistServiceImpl {
	return &OuterSpecialistServiceImpl{
		repos: repos,
	}
}
