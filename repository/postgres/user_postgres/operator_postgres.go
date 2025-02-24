package userPostgres

import (
	"database/sql"
	"fmt"
	"main/domain/entities"
	"main/repository/postgres"
	persistance "main/repository/postgres/entities_models"
	persistanceMappers "main/repository/postgres/mappers"
	"main/service/repository"
)

type OperatorPostgres struct {
	repository.OperatorRepository
	db *sql.DB
}

func NewOperatorPostgres(db *sql.DB) *OperatorPostgres {
	return &OperatorPostgres{db: db}
}

func (repos *OperatorPostgres) GetOperationsList(pagination int) ([]entities.Transfer, error) {
	offset := (pagination - 1) * postgres.TransfersPerRequestLimit
	query := fmt.Sprintf("SELECT user_id, sender_acc_num, receiver_acc_num, amount FROM %s LIMIT %s OFFSET $1", postgres.TransfersTable, fmt.Sprint(postgres.TransfersPerRequestLimit))
	rows, err := repos.db.Query(query, fmt.Sprint(offset))
	if err != nil {
		return nil, err
	}
	transfersPersistanceList := make([]persistance.TransferPersistance, postgres.TransfersPerRequestLimit)
	transfersList := make([]entities.Transfer, postgres.TransfersPerRequestLimit)

	err = rows.Scan(&transfersPersistanceList[0].TransferOwnerId, &transfersPersistanceList[0].SenderAccountNum,
		&transfersPersistanceList[0].ReceiverAccountNum, &transfersPersistanceList[0].SumOfTransfer)
	if err != nil {
		return nil, err
	}
	tempTransfer, err := persistanceMappers.ToTransferEntitity(transfersPersistanceList[0], entities.NewZeroChecker()) 
	transfersList[0] = tempTransfer
	if err != nil {
		return nil, err
	}
	i := 1
	for rows.Next() {
		err = rows.Scan(&transfersPersistanceList[i].TransferOwnerId, &transfersPersistanceList[i].SenderAccountNum,
			&transfersPersistanceList[i].ReceiverAccountNum, &transfersPersistanceList[i].SumOfTransfer)
		if err != nil {
			return nil, err
		}
		tempTransfer, err := persistanceMappers.ToTransferEntitity(transfersPersistanceList[i], entities.NewZeroChecker()) 
		transfersList[i] = tempTransfer
		if err != nil {
			return nil, err
		}
		i++
	}
	return transfersList, nil
}

func (repos *OperatorPostgres) CancelTransferOperation(operationId int) error {
	//DO THE WHOLE THING
	return nil
}

func (repos *OperatorPostgres) ApprovePaymentRequest(requestId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_approved=true WHERE request_id=$1", postgres.PaymentRequestsTable)
	_, err := repos.db.Exec(query, requestId)
	if err != nil {
		return err
	}
	return nil
}