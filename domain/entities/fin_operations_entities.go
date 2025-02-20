package entities

import (
	"errors"
	domainErrors "main/domain/entities/domain_errors"

	"github.com/shopspring/decimal"
)

type MoneyAmount = int64
type BankName = string
type BankIdentificationNum = string
type CreditRate = decimal.Decimal
type DayLength = int64

type BankAccount struct {
	amount                    MoneyAmount
	accountIdenitificationNum AccountIdenitificationNum
	bankFullName              BankName
	bankIdentificationNum     BankIdentificationNum
	status                    Status
	outsideInfo               BankAccountOutside
}

type BankAccountOutside interface {
	doesBankExist(bankIdentifNum BankIdentificationNum, bankFullName BankName) (bool, error)
}

type Status struct {
	isBlocked bool
	isFrozen  bool
}

type InstallmentPlan struct {
	info     Credit
	duration int
}

type PaymentRequest struct {
	amount   int
	clientId int
}

type Credit struct {
	bankProviderName          BankName
	accountIdenitificationNum AccountIdenitificationNum
	rate                      CreditRate
	principle                 MoneyAmount
	period                    DayLength
}

type Transfer struct {
	transferOwnerId    int
	SenderAccountNum   AccountIdenitificationNum
	SumOfTransfer      MoneyAmount
	ReceiverAccountNum AccountIdenitificationNum
	outsideInfo        TransferOutside
}

type TransferOutside interface {
	doesAccountBelongTo(accountNum AccountIdenitificationNum, userId int) (bool, error)
	accountMoneyAmount(accountNum AccountIdenitificationNum, userId int) (MoneyAmount, error)
}

func (trasnfer *Transfer) ValidateTransfer() error {
	return errors.Join(
		trasnfer.ValidateMoneyAmount(),
		trasnfer.ValidateAccounts(),
	)
}

func (transfer *Transfer) ValidateAccounts() error {
	senderBelonging, err := transfer.outsideInfo.doesAccountBelongTo(transfer.SenderAccountNum, transfer.transferOwnerId)
	if !senderBelonging {
		return domainErrors.NewNotPermitted("sender account does not belong to user")
	}
	if err != nil {
		return err
	}
	receiverBelonging, err := transfer.outsideInfo.doesAccountBelongTo(transfer.SenderAccountNum, transfer.transferOwnerId)
	if !receiverBelonging {
		return domainErrors.NewNotPermitted("receiver account does not belong to user")
	}
	if err != nil {
		return err
	}
	return nil
}

func (transfer *Transfer) ValidateMoneyAmount() error {
	if transfer.SumOfTransfer < 0 {
		return domainErrors.NewInvalidField("invalid money amount")
	}
	Money, err := transfer.outsideInfo.accountMoneyAmount(transfer.SenderAccountNum, transfer.transferOwnerId)
	if err != nil {
		return err
	}
	if Money < transfer.SumOfTransfer {
		return domainErrors.NewInvalidField("not enough money for a transfer")
	}
	return nil
}

func NewBankAccount(amount MoneyAmount, accountIdentifNum AccountIdenitificationNum, bankFullName BankName, bankIdentifNum BankIdentificationNum) (*BankAccount, error) {
	account := &BankAccount{
		amount:                    amount,
		accountIdenitificationNum: accountIdentifNum,
		bankFullName:              bankFullName,
		bankIdentificationNum:     bankIdentifNum,
		status:                    Status{},
	}
	if err := account.ValidateBankAccount(); err != nil {
		return nil, err
	}
	return account, nil
}

func (account *BankAccount) ValidateBankAccount() error {
	err := errors.Join(
		account.ValidateMoneyAmount(),
		account.ValidateStatus(),
		account.ValidateFullName(),
		account.ValidateBank(),
	)
	return err
}

func (account *BankAccount) ValidateMoneyAmount() error {
	if account.MoneyAmount() < 0 {
		return domainErrors.NewInvalidField("invalid money amount")
	}
	return nil
}

func (account *BankAccount) ValidateStatus() error {
	if account.status.isBlocked && account.status.isFrozen {
		return domainErrors.NewInvalidField("invalid account statuses")
	}
	return nil
}

func (account *BankAccount) ValidateFullName() error {
	if len(account.bankFullName) == 0 {
		return domainErrors.NewInvalidField("invalid bank full name")
	}
	return nil
}

func (account *BankAccount) ValidateBank() error {
	exists, err := account.outsideInfo.doesBankExist(account.bankIdentificationNum, account.bankFullName)
	if !exists {
		return domainErrors.NewInvalidField("such bank doesn't exist")
	}
	return err
}

func (account *BankAccount) MoneyAmount() MoneyAmount {
	return account.amount
}

func (account *BankAccount) AccountIdenitificationNum() AccountIdenitificationNum {
	return account.accountIdenitificationNum
}

func (account *BankAccount) BankFullName() BankName {
	return account.bankFullName
}

func (account *BankAccount) BankIdentificationNum() BankIdentificationNum {
	return account.bankIdentificationNum
}

func (account BankAccount) IsFrozen() bool {
	return account.status.isFrozen
}

func (account BankAccount) IsBlocked() bool {
	return account.status.isBlocked
}

func (*Transfer) NewTransfer(transferOwnerId int, SenderNum AccountIdenitificationNum, ReceiverNum AccountIdenitificationNum, Amount MoneyAmount) (*Transfer, error) {
	transferValue := &Transfer{
		transferOwnerId:    transferOwnerId,
		SenderAccountNum:   SenderNum,
		ReceiverAccountNum: ReceiverNum,
		SumOfTransfer:      Amount,
	}
	err := transferValue.ValidateTransfer()
	if err != nil {
		return nil, err
	}
	return transferValue, nil
}
