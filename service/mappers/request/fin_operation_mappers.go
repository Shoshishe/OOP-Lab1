package request_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/request"
)

func ToLoanEntity(model request.LoanModel) (*entities.Loan, error) {
	entity, err := entities.NewLoan(
		model.BankProviderName, model.AccountIdentifNum,
		model.Rate, model.LoanAmount, model.EndOfLoanTerm,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func ToTransferEntitity(model request.TransferModel) (*entities.Transfer, error) {
	entity, err := entities.NewTransfer(
		model.TransferOwnerId, model.SenderAccountNum, 
		model.ReceiverAccountNum, model.SumOfTransfer,
	)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
