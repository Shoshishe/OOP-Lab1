package request_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

func ToLoanEntity(model request.LoanModel) (*entities.Loan, error) {
	entity, err := entities.NewLoan(
		model.BankProviderName, model.AccountIdentifNum,
		model.Rate, model.LoanAmount, model.StartOfLoanTerm, model.EndOfLoanTerm,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func ToTransferEntitity(model request.TransferModel, validator entities.Validator) (*entities.Transfer, error) {
	entity, err := entities.NewTransfer(
		model.TransferOwnerId, model.SenderAccountNum,
		model.ReceiverAccountNum, model.SumOfTransfer,
		validator,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func ToInstallmentPlanEntity(model request.InstallmentPlanModel) (*entities.InstallmentPlan, error) {
	entity, err := entities.NewInstallmentPlan(
		model.BankProviderName, model.AmountForPayment,
		model.CountOfPayments, model.StartOfTerm, model.EndOfTerm,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func ToRequestEntity(model request.PaymentRequestModel) (*entities.PaymentRequest, error) {
	entity, err := entities.NewPaymentRequest(
		entities.MoneyAmount(model.Amount), model.AccountNum, 
		model.ClientId, model.CompanyId,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
