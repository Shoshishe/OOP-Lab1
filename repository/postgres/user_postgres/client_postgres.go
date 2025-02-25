package userPostgres

import (
	"database/sql"
	"fmt"
	"main/domain/entities"
	"main/repository/postgres"
	"main/service/repository"
	"time"
)

type ClientPostgres struct {
	repository.ClientRepository
	db *sql.DB
}

func (clientRepo *ClientPostgres) SendCreditsForPayment(req entities.PaymentRequest, usrId int) error {
	tx, err := clientRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (client_id, company_id, account_num, full_name, amount) VALUES ($1,$2,$3,$4,$5)", postgres.PaymentRequestsTable)
	_, err = clientRepo.db.Exec(query, req.ClientId(), req.CompanyId(), req.AccountNum(), req.RequesterFullName(), req.Amount())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 5)
	args = append(args,fmt.Sprint(req.ClientId()), fmt.Sprint(req.CompanyId()), req.AccountNum(), req.RequesterFullName(), fmt.Sprint(req.Amount()))
	err = postgres.InsertAction(tx, clientRepo.db, "SendCredits", args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (clientRepo *ClientPostgres) ReverseSendCreditsForPayment(req entities.PaymentRequest, usrId int) error {
	tx, err := clientRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE company_id=$1 AND client_id=$2", postgres.PaymentRequestsTable)
	_, err = clientRepo.db.Exec(query,req.CompanyId(),req.ClientId())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 5)
	args = append(args,fmt.Sprint(req.ClientId()), fmt.Sprint(req.CompanyId()), req.AccountNum(), req.RequesterFullName(), fmt.Sprint(req.Amount()))
	err = postgres.ReverseAction(tx, clientRepo.db, args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (clientRepo *ClientPostgres) TakeInstallmentPlan(plan entities.InstallmentPlan, usrId int) error {
	tx, err := clientRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (bank_provider_name, loan_amount, count_of_payments,start_date,end_date) VALUES ($1,$2,$3,$4,$5)", postgres.LoansTable)
	_, err = clientRepo.db.Exec(query, plan.BankProviderName(), plan.AmountForPayment(), plan.CountOfPayments(), plan.StartOfTerm(), plan.EndOfTerm())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 5)
	args = append(args, plan.BankProviderName(), fmt.Sprint(plan.AmountForPayment()), fmt.Sprint(plan.CountOfPayments()), plan.StartOfTerm().Format(time.DateOnly), plan.EndOfTerm().Format(time.DateOnly))
	err = postgres.InsertAction(tx, clientRepo.db, "TakeInstallmentPlan", args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (clientRepo *ClientPostgres) ReverseTakeInstallmentPlan(plan entities.InstallmentPlan, usrId int) error {
	tx, err := clientRepo.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (bank_provider_name, loan_amount, count_of_payments,start_date,end_date) VALUES ($1,$2,$3,$4,$5)", postgres.LoansTable)
	_, err = clientRepo.db.Exec(query, plan.BankProviderName(), plan.AmountForPayment(), plan.CountOfPayments(), plan.StartOfTerm(), plan.EndOfTerm())
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 5)
	args = append(args, plan.BankProviderName(), fmt.Sprint(plan.AmountForPayment()), fmt.Sprint(plan.CountOfPayments()), plan.StartOfTerm().Format(time.DateOnly), plan.EndOfTerm().Format(time.DateOnly))
	postgres.ReverseAction(tx, clientRepo.db, args, usrId)
	tx.Commit()
	return nil
}

func (clientRepo *ClientPostgres) TakeLoan(loan entities.Loan, usrId int) error {
	tx, err := clientRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (bank_provider_name, account_identif_num, rate, loan_amount, start_of_loan, end_of_loan) VALUES ($1,$2,$3,$4,$5,$6)", postgres.LoansTable)
	_, err = clientRepo.db.Exec(query, loan.BankProviderName(), loan.AccountIdenitificationNum(), loan.Rate(), loan.LoanAmount(),
		loan.StartOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly), loan.EndOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly))
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, loan.BankProviderName(), loan.AccountIdenitificationNum(), loan.Rate(), fmt.Sprint(loan.LoanAmount()),
		loan.StartOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly), loan.EndOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly))
	err = postgres.InsertAction(tx, clientRepo.db, "TakeLoan", args, usrId)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (clientRepo *ClientPostgres) ReverseTakeLoan(loan entities.Loan, usrId int) error {
	tx, err := clientRepo.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf("INSERT INTO %s (bank_provider_name, account_identif_num, rate, loan_amount, start_of_loan, end_of_loan) VALUES ($1,$2,$3,$4,$5,$6)", postgres.LoansTable)
	_, err = clientRepo.db.Exec(query, loan.BankProviderName(), loan.AccountIdenitificationNum(), loan.Rate(), loan.LoanAmount(),
		loan.StartOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly), loan.EndOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly))
	if err != nil {
		tx.Rollback()
		return err
	}
	args := make([]string, 0, 3)
	args = append(args, loan.BankProviderName(), loan.AccountIdenitificationNum(), loan.Rate(), fmt.Sprint(loan.LoanAmount()),
		loan.StartOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly), loan.EndOfLoanTerm().Truncate(24*time.Hour).Format(time.DateOnly))
	postgres.ReverseAction(tx, clientRepo.db, args, usrId)
	tx.Commit()
	return nil
}

func NewClientPostgres(db *sql.DB) *ClientPostgres {
	return &ClientPostgres{
		db: db,
	}
}
