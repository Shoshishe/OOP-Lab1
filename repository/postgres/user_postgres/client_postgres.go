package userPostgres

import (
	"database/sql"
	"fmt"
	"main/domain/entities"
	"main/repository/postgres"
	"main/service"
	"time"
)

type ClientPostgres struct {
	service.ClientRepository
	db *sql.DB
}

func (clientRepo *ClientPostgres) TakeInstallmentPlan(plan entities.InstallmentPlan) error {
	query := fmt.Sprintf("INSERT INTO %s (bank_provider_name, loan_amount, count_of_payments,start_date,end_date) VALUES ($1,$2,$3,$4,$5)", postgres.LoansTable)
	_, err := clientRepo.db.Exec(query, plan.BankProviderName(), plan.AmountForPayment(), plan.CountOfPayments(), plan.StartOfTerm(), plan.EndOfTerm())
	if err != nil {
		return err
	}
	return nil
}

func (clientRepo *ClientPostgres) TakeLoan(loan entities.Loan) error {
	query := fmt.Sprintf("INSERT INTO %s (bank_provider_name, account_identif_num, rate, loan_amount, start_of_loan, end_of_loan) VALUES ($1,$2,$3,$4,$5,$6)", postgres.LoansTable)
	_, err := clientRepo.db.Exec(query, loan.BankProviderName(), loan.AccountIdenitificationNum(), loan.Rate(), loan.LoanAmount(), loan.StartOfLoanTerm().Truncate(24*time.Hour), loan.EndOfLoanTerm().Truncate(24*time.Hour))
	if err != nil {
		return err
	}
	return nil
}

func NewClientPostgres(db *sql.DB) *ClientPostgres {
	return &ClientPostgres{
		db: db,
	}
}
