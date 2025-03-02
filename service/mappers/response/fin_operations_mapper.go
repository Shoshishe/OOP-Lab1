package response_mappers

import (
	"main/domain/entities"
	"main/service/entities_models/response"
)

func ToTransferModel(transfer *entities.Transfer) *response.TransferModel {
	return &response.TransferModel{
		SenderAccountNum: transfer.SenderAccountNum(),
		ReceiverAccountNum: transfer.ReceiverAccountNum(),
		SumOfTransfer: int64(transfer.SumOfTransfer()),
	}
}