package service

import (
	"main/domain/entities"
	domainErrors "main/domain/entities/domain_errors"
	serviceErrors "main/service/errors"
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
	"strconv"
	"time"
)

const ()

type Reverser struct {
	serviceInterfaces.Reverser
	accountReverser         repository.AccountReverserRepository
	clientReverser          repository.ClientActionsReverserRepository
	bankReverser            repository.BankActionReverserRepository
	operatorReverser        repository.OperatorActionsReverserRepository
	outerSpecialistReverser repository.OuterSpecialistReverserRepository
	infoRepos               repository.ReverserInfoRepository
}

//TODO

func NewReverser(accountReverser repository.AccountReverserRepository, clientReverser repository.ClientActionsReverserRepository,
	operatorReverser repository.OperatorActionsReverserRepository,
	infoRepos repository.ReverserInfoRepository) *Reverser {
	return &Reverser{
		accountReverser:  accountReverser,
		clientReverser:   clientReverser,
		operatorReverser: operatorReverser,
		infoRepos:        infoRepos,
	}
}

func (reverser *Reverser) getAction(actionId int) (string, []string, error) {
	action, err := reverser.infoRepos.GetAction(actionId)
	if err != nil {
		return "", nil, err
	}
	return action.ActionName(), action.ActionArgs(), nil
}

func (reverser *Reverser) Reverse(actionId int, usrId int, usrRole int) error {
	actionName, args, err := reverser.getAction(actionId)
	if err != nil {
		return err
	}
	switch actionName {
	case repository.AccountCreationAction:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		if len(args) < 3 {
			return domainErrors.NewInvalidField("invalid action args")
		}
		bankAccount, err := entities.NewBankAccount(args[0], args[1], args[2])
		if err != nil {
			return err
		}
		return reverser.accountReverser.ReverseAccountCreation(*bankAccount, usrId)
	case repository.BlockAccountAction:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		if len(args) < 1 {
			return domainErrors.NewInvalidField("invalid action args")
		}
		return reverser.accountReverser.ReverseAccountBlock(args[0], usrId)
	case repository.FreezeAccountAction:
		if len(args) < 3 {
			return domainErrors.NewInvalidField("invalid action args")
		}
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		return reverser.accountReverser.ReverseAccountFreeze(args[0], usrId)
	case repository.TransferMoneyAction:
		if len(args) < 3 {
			return domainErrors.NewInvalidField("invalid args counter")
		}
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted too reverse such action")
		}
		amount, err := strconv.Atoi(args[2])
		if err != nil {
			return domainErrors.NewInvalidField("invalid arg type")
		}
		transfer, err := entities.NewTransfer(0, args[0], args[1], entities.MoneyAmount(amount), entities.NewZeroChecker())
		if err != nil {
			return err
		}
		return reverser.accountReverser.ReverseMoneyTransfer(*transfer, usrId)
	case repository.AddBankAction:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		if len(args) < 6 {
			return domainErrors.NewInvalidField("invalid args count")
		}
		bankInfo, err := entities.NewCompany(args[0], args[1], args[2], args[3], args[4])
		if err != nil {
			return err
		}
		bankType := args[5]
		bank, err := entities.NewBank(*bankInfo, bankType)
		if err != nil {
			return err
		}
		return reverser.bankReverser.ReverseBankAddition(*bank, usrId)
	case repository.SendPaymentRequest:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		if len(args) < 5 {
			return domainErrors.NewInvalidField("invalid args count")
		}
		amount, err := strconv.Atoi(args[4])
		if err != nil {
			return err
		}
		clientId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		companyId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		req, err := entities.NewPaymentRequest(entities.MoneyAmount(amount), args[2], args[3], clientId, companyId)
		if err != nil {
			return err
		}
		return reverser.clientReverser.ReverseSendCreditsForPayment(*req, usrId)
	case repository.TakeInstallmentPlanAction:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		amount, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		countOfPayments, err := strconv.Atoi(args[2])
		if err != nil {
			return err
		}
		if len(args) < 6 {
			return domainErrors.NewInvalidField("invalid args count")
		}
		start_date, err := time.Parse(time.DateOnly, args[3])
		if err != nil {
			return err
		}
		end_date, err := time.Parse(time.DateOnly, args[4])
		if err != nil {
			return err
		}
		loan, err := entities.NewInstallmentPlan(args[0], entities.MoneyAmount(amount), entities.Count(countOfPayments), start_date, end_date, args[5])
		if err != nil {
			return err
		}
		return reverser.clientReverser.ReverseTakeInstallmentPlan(*loan, usrId)
	case repository.TakeLoanAction:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		if len(args) < 5 {
			return domainErrors.NewInvalidField("invalid args count")
		}
		amount, err := strconv.Atoi(args[3])
		if err != nil {
			return err
		}
		start_date, err := time.Parse(time.DateOnly, args[4])
		if err != nil {
			return err
		}
		end_date, err := time.Parse(time.DateOnly, args[5])
		if err != nil {
			return err
		}
		loan, err := entities.NewLoan(args[0], args[1], args[2], entities.MoneyAmount(amount), start_date, end_date)
		if err != nil {
			return err
		}
		return reverser.clientReverser.ReverseTakeLoan(*loan, usrId)
	case repository.CancelTransferAction:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		operationId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		return reverser.operatorReverser.ReverseCancelTransferOperation(operationId, usrId)
	case repository.TransferRequestAction:
		if usrRole != entities.RoleAdmin {
			return serviceErrors.NewRoleError("not permitted to reverse such action")
		}
		ownerId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		sumOfTransfer, err := strconv.Atoi(args[3])
		if err != nil {
			return err
		}
		transfer, err := entities.NewTransfer(ownerId, args[1], args[2], entities.MoneyAmount(sumOfTransfer), entities.NewZeroChecker())
		if err != nil {
			return err
		}
		return reverser.outerSpecialistReverser.ReverseTransferRequest(*transfer, usrId)
	}
	return nil
}
