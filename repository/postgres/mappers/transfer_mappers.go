package persistanceMappers

import (
	"main/domain/entities"
	persistance "main/repository/postgres/entities_models"
)

func ToTransferEntitity(transfer persistance.TransferPersistance, validator entities.Validator) (entities.Transfer, error) {
	entity, err := entities.NewTransfer(
		transfer.TransferOwnerId, transfer.SenderAccountNum,
		transfer.ReceiverAccountNum, entities.MoneyAmount(transfer.SumOfTransfer),
		validator,
	)
	if err != nil {
		return entities.Transfer{}, err
	}
	return *entity, nil
}
