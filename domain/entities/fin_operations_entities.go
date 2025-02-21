package entities

import (
	"errors"
	domainErrors "main/domain/entities/domain_errors"
	"time"

	"github.com/shopspring/decimal"
)

type MoneyAmount = int64
type BankName = string
type BankIdentificationNum = string
type CreditRate = decimal.Decimal
type Date = time.Time

type InstallmentPlan struct {
	bankProviderName BankName
	endOfTerm        Date
}

type PaymentRequest struct {
	amount   int
	clientId int
}

var threeMonths, _ = time.Parse(time.DateOnly, "0000-04-01")
var sixMonths, _ = time.Parse(time.DateOnly, "0000-07-01")
var twelveMonths, _ = time.Parse(time.DateOnly, "0001-01-01")
var twentyFourMonths, _ = time.Parse(time.DateOnly, "0002-01-01")

func AddDates(currentDate, addedValue time.Time) time.Time {
	return currentDate.AddDate(addedValue.Year(), int(addedValue.Month())-1, addedValue.Day()-1)
}

type LoanOutside interface {
	CheckIfLoanExists(loanId int) (bool, error)
}
type Loan struct {
	loanId                    int
	bankProviderName          BankName
	accountIdenitificationNum AccountIdenitificationNum
	rate                      CreditRate
	loanAmount                MoneyAmount
	endOfLoanTerm             Date
	isAccepted				  bool
	outsideInfo               LoanOutside
}

func (loan *Loan) LoanId() int {
	return loan.loanId
}

func (loan *Loan) BankProviderName() BankName {
	return loan.accountIdenitificationNum
}

func (loan *Loan) AccountIdenitificationNum() AccountIdenitificationNum {
	return loan.accountIdenitificationNum
}

func (loan *Loan) Rate() CreditRate {
	return loan.rate
}

func (loan *Loan) LoanAmount() MoneyAmount {
	return loan.loanAmount
}

func (loan *Loan) EndOfLoanTerm() Date {
	return loan.endOfLoanTerm
}

func (loan *Loan) IsAccepted() bool {
	return loan.isAccepted
}

func NewLoan(bankProviderName BankName, accountNum AccountIdenitificationNum, rate CreditRate, loanAmount MoneyAmount, endOfLoanTerm Date) (*Loan, error) {
	loan := &Loan{
		bankProviderName:          bankProviderName,
		accountIdenitificationNum: accountNum,
		rate:                      rate,
		loanAmount:                loanAmount,
		endOfLoanTerm:             endOfLoanTerm,
	}
	err := loan.ValidateLoan()
	if err != nil {
		return nil, err
	}
	return loan, nil
}

func (loan *Loan) ValidateLoan() error {
	err := errors.Join(
		loan.ValidateAccountIdentifNum(),
		loan.ValidateBankProviderName(),
		loan.ValidateEndOfLoanTerm(),
		loan.ValidateLoanAmount(),
		loan.ValidateRate(),
	)
	return err
}

func (loan *Loan) ValidateBankProviderName() error {
	var err error
	if len(loan.bankProviderName) == 0 {
		err = domainErrors.NewInvalidField("empty bank provider name")
	}
	return err
}

func (loan *Loan) ValidateAccountIdentifNum() error {
	var err error
	if len(loan.accountIdenitificationNum) == 0 {
		err = domainErrors.NewInvalidField("empty account identif num")
	}
	return err
}

func (loan *Loan) ValidateRate() error {
	var err error
	if loan.rate.LessThan(decimal.New(1, 1)) {
		err = domainErrors.NewInvalidField("invalid credit rate")
	}
	return err
}

func (loan *Loan) ValidateLoanAmount() error {
	var err error
	if loan.loanAmount <= 0 {
		err = domainErrors.NewInvalidField("invalid loan amount")
	}
	return err
}

func (loan *Loan) ValidateEndOfLoanTerm() error {
	exists, err := loan.outsideInfo.CheckIfLoanExists(loan.loanId)
	if err != nil {
		return err
	}
	if loan.endOfLoanTerm.Before(time.Now().Truncate(time.Hour * 24)) {
		if !exists {
			switch loan.endOfLoanTerm {
			case AddDates(time.Now().Truncate(time.Hour*24), threeMonths):
			case AddDates(time.Now().Truncate(time.Hour*24), sixMonths):
			case AddDates(time.Now().Truncate(time.Hour*24), twelveMonths):
			case AddDates(time.Now().Truncate(time.Hour*24), twentyFourMonths):
			default:
				if !loan.endOfLoanTerm.After(AddDates(time.Now().Truncate(time.Hour*24), twentyFourMonths)) {
					return domainErrors.NewInvalidField("unacceptable loan term")
				}
			}
		} else {
			err = domainErrors.NewInvalidField("end of loan is before the current time")
		}
	}
	return err
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
