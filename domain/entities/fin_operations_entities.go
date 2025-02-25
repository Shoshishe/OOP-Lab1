package entities

import (
	"errors"
	domainErrors "main/domain/entities/domain_errors"
	"strings"
	"time"
)

type MoneyAmount = int64
type BankName = string
type BankIdentificationNum = string
type CreditRate = string
type Date = time.Time
type Count = int16

type Action struct {
	actionId   int
	actionName string
	actionArgs []string
}

func (action *Action) ActionId() int {
	return action.actionId
}

func (action *Action) ActionName() string {
	return action.actionName
}

func (action *Action) ActionArgs() []string {
	return action.actionArgs
}

func NewAction(actionId int, actionName string, actionArgs []string) *Action {
	return &Action{
		actionId:   actionId,
		actionName: actionName,
		actionArgs: actionArgs,
	}
}

type PaymentRequest struct {
	amount            MoneyAmount
	accountNum        AccountIdenitificationNum
	requesterFullName FullName
	clientId          int
	companyId         int
}

func (req *PaymentRequest) RequesterFullName() FullName {
	return req.requesterFullName
}

func (req *PaymentRequest) Amount() MoneyAmount {
	return req.amount
}

func (req *PaymentRequest) AccountNum() AccountIdenitificationNum {
	return req.accountNum
}

func (req *PaymentRequest) ClientId() int {
	return req.clientId
}

func (req *PaymentRequest) CompanyId() int {
	return req.companyId
}

func NewPaymentRequest(amount MoneyAmount, accountNum string, requsterFullName string, clientId int, companyId int) (*PaymentRequest, error) {
	req := &PaymentRequest{
		amount:     amount,
		accountNum: accountNum,
		clientId:   clientId,
		companyId:  companyId,
	}
	err := req.ValidateRequest()
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (req *PaymentRequest) ValidateRequest() error {
	return errors.Join(
		req.ValidateAccountNum(),
		req.ValidateAmount(),
		req.ValidateId(),
	)
}

func (req *PaymentRequest) ValidateAmount() error {
	if req.amount <= 0 {
		return domainErrors.NewInvalidField("invalid requested amount")
	}
	return nil
}

func (req *PaymentRequest) ValidateAccountNum() error {
	if len(req.accountNum) == 0 {
		return domainErrors.NewInvalidField("invalid account number")
	}
	return nil
}

func (req *PaymentRequest) ValidateId() error {
	if req.clientId <= 0 || req.companyId <= 0 {
		return domainErrors.NewInvalidField("incorrect user id")
	}
	return nil
}

type InstallmentPlan struct {
	bankProviderName  BankName
	amountForPayment  MoneyAmount
	accountIdentifNum AccountIdenitificationNum
	countOfPayments   Count
	startOfTerm       Date
	endOfTerm         Date
	isAccepted        bool
}

func NewInstallmentPlan(bankProviderName BankName, amountForPayment MoneyAmount, countOfPayments Count, startOfTerm Date, endOfTerm Date, accountIdentifNum AccountIdenitificationNum) (*InstallmentPlan, error) {
	planEntity := &InstallmentPlan{
		bankProviderName:  bankProviderName,
		amountForPayment:  amountForPayment,
		countOfPayments:   countOfPayments,
		startOfTerm:       startOfTerm,
		endOfTerm:         endOfTerm,
		accountIdentifNum: accountIdentifNum,
	}
	err := planEntity.ValidatePlan()
	if err != nil {
		return nil, err
	}
	return planEntity, nil
}

func (plan *InstallmentPlan) ValidatePlan() error {
	err := errors.Join(
		plan.ValidateBankProviderName(),
		plan.ValidateCountOfPayments(),
		plan.ValidateTermDates(),
		plan.ValidateMoneyAmount(),
	)
	return err
}

func (plan *InstallmentPlan) ValidateTermDates() error {
	if plan.startOfTerm.Truncate(24 * time.Hour).Equal(plan.endOfTerm.Truncate(24 * time.Hour)) {
		return domainErrors.NewInvalidField("term start and end are the same")
	}
	if plan.startOfTerm.Truncate(24 * time.Hour).After(plan.endOfTerm.Truncate(24 * time.Hour)) {
		return domainErrors.NewInvalidField("term start is after the term end")
	}
	return nil
}

func (plan *InstallmentPlan) ValidateCountOfPayments() error {
	if plan.countOfPayments < 1 {
		return domainErrors.NewInvalidField("invalid count of payments")
	}
	return nil
}

func (plan *InstallmentPlan) ValidateBankProviderName() error {
	if len(plan.bankProviderName) == 0 {
		return domainErrors.NewInvalidField("empty bank provider name")
	}
	return nil
}

func (plan *InstallmentPlan) ValidateMoneyAmount() error {
	if plan.amountForPayment <= 0 {
		return domainErrors.NewInvalidField("invalid payment amount")
	}
	return nil
}

func (plan *InstallmentPlan) BankProviderName() string {
	return plan.bankProviderName
}

func (plan *InstallmentPlan) AmountForPayment() int {
	return int(plan.amountForPayment)
}

func (plan *InstallmentPlan) AccountIdentifNum() string {
	return plan.accountIdentifNum
}

func (plan *InstallmentPlan) CountOfPayments() int16 {
	return plan.countOfPayments
}

func (plan *InstallmentPlan) StartOfTerm() Date {
	return plan.startOfTerm
}

func (plan *InstallmentPlan) EndOfTerm() Date {
	return plan.endOfTerm
}

func (plan *InstallmentPlan) IsAccepted() bool {
	return plan.isAccepted
}

type Loan struct {
	bankProviderName          BankName
	accountIdenitificationNum AccountIdenitificationNum
	rate                      CreditRate
	loanAmount                MoneyAmount
	startOfLoanTerm           Date
	endOfLoanTerm             Date
	isAccepted                bool
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

func (loan *Loan) StartOfLoanTerm() Date {
	return loan.startOfLoanTerm
}

func (loan *Loan) EndOfLoanTerm() Date {
	return loan.endOfLoanTerm
}

func (loan *Loan) IsAccepted() bool {
	return loan.isAccepted
}

func NewLoan(bankProviderName BankName, accountNum AccountIdenitificationNum, rate CreditRate, loanAmount MoneyAmount, startOfLoanTerm Date, endOfLoanTerm Date) (*Loan, error) {
	loan := &Loan{
		bankProviderName:          bankProviderName,
		accountIdenitificationNum: accountNum,
		rate:                      rate,
		loanAmount:                loanAmount,
		startOfLoanTerm:           startOfLoanTerm,
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
		loan.ValidateLoanTerm(),
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
	if loan.rate == "" || !strings.Contains(loan.rate, "%") {
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

func isThreeMonthsApart(startDate time.Time, endDate time.Time) bool {
	if startDate.Day() == endDate.Day() {
		if startDate.Year() == endDate.Year() {
			if endDate.Month()-startDate.Month() == 3 {
				return true
			}
		}
		return false
	}
	return false
}

func isSixMonthsApart(startDate time.Time, endDate time.Time) bool {
	if startDate.Day() == endDate.Day() {
		if startDate.Year() == endDate.Year() {
			if endDate.Month()-startDate.Month() == 6 {
				return true
			}
		}
	}
	return false
}

func isYearApart(startDate time.Time, endDate time.Time) bool {
	if startDate.Day() == endDate.Day() {
		if startDate.Month() == endDate.Month() {
			if endDate.Year()-startDate.Year() == 1 {
				return true
			}
		}
	}
	return false
}

func isTwoYearsApart(startDate time.Time, endDate time.Time) bool {
	if startDate.Day() == endDate.Day() {
		if startDate.Month() == endDate.Month() {
			if endDate.Year()-startDate.Year() == 2 {
				return true
			}
		}
	}
	return false
}

func IsMoreThanTwoYearsApart(startDate time.Time, endDate time.Time) bool {
	if endDate.Year()-startDate.Year() == 2 {
		if endDate.Month() > startDate.Month() {
			return true
		} else if endDate.Month() < startDate.Month() {
			return false
		} else {
			if endDate.Day() > startDate.Day() {
				return true
			} else {
				return false
			}
		}
	} else if endDate.Year()-startDate.Year() > 2 {
		return true
	}
	return false
}

func (loan *Loan) ValidateLoanTerm() error {
	if !(isThreeMonthsApart(loan.startOfLoanTerm, loan.endOfLoanTerm) || isSixMonthsApart(loan.startOfLoanTerm, loan.endOfLoanTerm) ||
		isYearApart(loan.startOfLoanTerm, loan.endOfLoanTerm) || isTwoYearsApart(loan.startOfLoanTerm, loan.endOfLoanTerm) ||
		IsMoreThanTwoYearsApart(loan.startOfLoanTerm, loan.endOfLoanTerm)) {
		return domainErrors.NewInvalidField("invalid loan terms")
	}
	return nil
}

type Validator interface {
	ValidateAccount(transfer *Transfer) error
	ValidateMoneyAmount(transfer *Transfer) error
}

type UserTransferOutside interface {
	DoesAccountBelongToUser(accountNum AccountIdenitificationNum, userId int) (bool, error)
	DoesAccountExist(AccountIdenitificationNum) (bool, error)
	AccountMoneyAmount(accountNum AccountIdenitificationNum) (MoneyAmount, error)
}
type userAccountChecker struct {
	Validator
	outsideInfo UserTransferOutside
}

func (checker *userAccountChecker) ValidateAccount(transfer *Transfer) error {
	if transfer.senderAccountNum == transfer.receiverAccountNum {
		return domainErrors.NewInvalidField("sender account is equal to receiver account")
	}
	senderBelonging, err := checker.outsideInfo.DoesAccountBelongToUser(transfer.SenderAccountNum(), transfer.TransferOwnerId())
	if err != nil {
		return err
	}
	if !senderBelonging {
		return domainErrors.NewInvalidField("account does not belong to sender")
	}
	receiverExists, err := checker.outsideInfo.DoesAccountExist(transfer.ReceiverAccountNum())
	if err != nil {
		return err
	}
	if !receiverExists {
		return domainErrors.NewInvalidField("receiver acc does not exist")
	}
	return nil
}

func NewUserAccountChecker(outsideInfo UserTransferOutside) *userAccountChecker {
	return &userAccountChecker{outsideInfo: outsideInfo}
}

type CompanyTransferOutside interface {
	DoesAccountBelongToOuterCompany(accountNum AccountIdenitificationNum, specialistId int) (bool, error)
	DoesAccountBelongToNonOuterUser(accountNum AccountIdenitificationNum, specialistId int) (bool, error)
	DoesAccountBelongToUser(AccountIdenitificationNum, int) (bool, error)
	AccountMoneyAmount(accountNum AccountIdenitificationNum) (MoneyAmount, error)
}

type companyAccountChecker struct {
	Validator
	outsideInfo CompanyTransferOutside
}

func (checker *companyAccountChecker) ValidateAccount(transfer *Transfer) error {
	if transfer.senderAccountNum == transfer.receiverAccountNum {
		return domainErrors.NewInvalidField("sender account is equal to receiver account")
	}
	belongsToSender, err := checker.outsideInfo.DoesAccountBelongToUser(transfer.SenderAccountNum(), transfer.transferOwnerId)
	if err != nil {
		return err
	}
	if !belongsToSender {
		return domainErrors.NewInvalidField("sender account does not belong to sender")
	}
	belongsToCompany, err := checker.outsideInfo.DoesAccountBelongToOuterCompany(transfer.SenderAccountNum(), transfer.TransferOwnerId())
	if err != nil {
		return err
	}
	if !belongsToCompany {
		belongsToOuterUser, err := checker.outsideInfo.DoesAccountBelongToNonOuterUser(transfer.SenderAccountNum(), transfer.TransferOwnerId())
		if err != nil {
			return err
		}
		if !belongsToOuterUser {
			return err
		}
	} else {
		return nil
	}
	return nil
}

func NewCompanyAccountChecker(outsideInfo CompanyTransferOutside) *companyAccountChecker {
	return &companyAccountChecker{outsideInfo: outsideInfo}
}

// PROCEED AT YOUR OWN RISK
type ZeroChecker struct {
	Validator
}

func (validator *ZeroChecker) ValidateAccount(transfer *Transfer) error {
	return nil
}

func (validator *ZeroChecker) ValidateMoneyAmount(transfer *Transfer) error {
	return nil
}

func NewZeroChecker() *ZeroChecker {
	return &ZeroChecker{}
}

type Transfer struct {
	transferOwnerId    int
	senderAccountNum   AccountIdenitificationNum
	sumOfTransfer      MoneyAmount
	receiverAccountNum AccountIdenitificationNum
	validator          Validator
}

func (trasnfer *Transfer) ValidateTransfer() error {
	return errors.Join(
		trasnfer.validator.ValidateMoneyAmount(trasnfer),
		trasnfer.validator.ValidateAccount(trasnfer),
	)
}

func (checker userAccountChecker) ValidateMoneyAmount(transfer *Transfer) error {
	if transfer.sumOfTransfer < 0 {
		return domainErrors.NewInvalidField("invalid money amount")
	}
	Money, err := checker.outsideInfo.AccountMoneyAmount(transfer.senderAccountNum)
	if err != nil {
		return err
	}
	if Money < transfer.sumOfTransfer {
		return domainErrors.NewInvalidField("not enough money for a transfer")
	}
	return nil
}

func (checker companyAccountChecker) ValidateMoneyAmount(transfer *Transfer) error {
	if transfer.sumOfTransfer < 0 {
		return domainErrors.NewInvalidField("invalid money amount")
	}
	Money, err := checker.outsideInfo.AccountMoneyAmount(transfer.senderAccountNum)
	if err != nil {
		return err
	}
	if Money < transfer.sumOfTransfer {
		return domainErrors.NewInvalidField("not enough money for a transfer")
	}
	return nil
}

func (transfer *Transfer) TransferOwnerId() int {
	return transfer.transferOwnerId
}

func (transfer *Transfer) SenderAccountNum() string {
	return transfer.senderAccountNum
}

func (transfer *Transfer) ReceiverAccountNum() string {
	return transfer.receiverAccountNum
}

func (transfer *Transfer) SumOfTransfer() int {
	return int(transfer.sumOfTransfer)
}

func NewTransfer(transferOwnerId int, SenderNum AccountIdenitificationNum, ReceiverNum AccountIdenitificationNum, Amount MoneyAmount, checker Validator) (*Transfer, error) {
	transferValue := &Transfer{
		transferOwnerId:    transferOwnerId,
		senderAccountNum:   SenderNum,
		receiverAccountNum: ReceiverNum,
		sumOfTransfer:      Amount,
		validator:          checker,
	}
	err := transferValue.ValidateTransfer()
	if err != nil {
		return nil, err
	}
	return transferValue, nil
}
