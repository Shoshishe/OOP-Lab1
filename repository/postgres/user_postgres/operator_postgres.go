package userPostgres

import (
	"database/sql"
	"errors"
	"fmt"
	"main/domain/entities"
	"main/repository/postgres"
	persistance "main/repository/postgres/entities_models"
	persistanceMappers "main/repository/postgres/mappers"
	"main/service/repository"
	"strconv"
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

func (repos *OperatorPostgres) CancelTransferOperation(operationId int, usrId int) error {
	var transferArgs []string
	var operationName string
	query := fmt.Sprintf("SELECT first_action_type,first_action_args FROM %s WHERE first_action_id=$1", postgres.ActionsTable)
	row := repos.db.QueryRow(query, operationId)
	err := row.Scan(&operationName, &transferArgs)
	if err == sql.ErrNoRows {
		query := fmt.Sprintf("SELECT second_action_type,second_action_args FROM %s WHERE second_action_id=$1", postgres.ActionsTable)
		row := repos.db.QueryRow(query, usrId)
		err := row.Scan(&operationName, &transferArgs)
		if err != nil {
			return err
		}
	}
	if operationName != "TransferMoney" {
		return errors.New("incorrect operation")
	}
	tx, err := repos.db.Begin()
	if err != nil {
		return err
	}
	senderAccountNum := transferArgs[0]
	receiverAccountNum := transferArgs[1]
	moneyAmount, err := strconv.Atoi(transferArgs[2])
	if err != nil {
		return err
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", postgres.AccountsTable)
	_, err = repos.db.Exec(changeSenderMoneyQuery, moneyAmount, senderAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", postgres.AccountsTable)
	_, err = repos.db.Exec(changeReceiverMoneyQuery, moneyAmount, receiverAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, senderAccountNum, receiverAccountNum, fmt.Sprint(moneyAmount))
	err = postgres.InsertAction(tx, repos.db, "CancelTransferMoney", args, usrId)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (repos *OperatorPostgres) ReverseCancelTransferOperation(operationId, usrId int) error {
	var transferArgs []string
	var operationName string
	query := fmt.Sprintf("SELECT first_action_type,first_action_args FROM %s WHERE first_action_id=$1", postgres.ActionsTable)
	row := repos.db.QueryRow(query, operationId)
	err := row.Scan(&operationName, &transferArgs)
	if err == sql.ErrNoRows {
		query := fmt.Sprintf("SELECT second_action_type,second_action_args FROM %s WHERE second_action_id=$1", postgres.ActionsTable)
		row := repos.db.QueryRow(query, usrId)
		err := row.Scan(&operationName, &transferArgs)
		if err != nil {
			return err
		}
	}
	if operationName != "TransferMoney" {
		return errors.New("incorrect operation")
	}
	tx, err := repos.db.Begin()
	if err != nil {
		return err
	}
	senderAccountNum := transferArgs[0]
	receiverAccountNum := transferArgs[1]
	moneyAmount, err := strconv.Atoi(transferArgs[2])
	if err != nil {
		return err
	}
	changeSenderMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount-$1 WHERE account_identif_num=$2", postgres.AccountsTable)
	_, err = repos.db.Exec(changeSenderMoneyQuery, moneyAmount, senderAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	changeReceiverMoneyQuery := fmt.Sprintf("UPDATE %s SET amount=amount+$1 WHERE account_identif_num=$2", postgres.AccountsTable)
	_, err = repos.db.Exec(changeReceiverMoneyQuery, moneyAmount, receiverAccountNum)
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, senderAccountNum, receiverAccountNum, fmt.Sprint(moneyAmount))
	err = postgres.ReverseAction(tx, repos.db, args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (repos *OperatorPostgres) ApprovePaymentRequest(requestId int, usrId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE request_id=$1", postgres.PaymentRequestsTable)
	_, err := repos.db.Exec(query, requestId)
	if err != nil {
		return err
	}
	return nil
}
