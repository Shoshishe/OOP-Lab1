package postgres

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

const (
	UsersTable = "users"
	BanksTable = "companies"
	AccountsTable = "accounts"
	LoansTable = "loans"
	PaymentRequestsTable = "requests"
	TransfersTable = "transfers"
	CompaniesTable = "companies"
	CompaniesWorkersTable = "companies_workers"
	ActionsTable = "actions"
	BanksPerRequestLimit = 50
	TransfersPerRequestLimit = 50
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	DbName   string
	Password string
	SSLMode  string
}

func NewPostgresDb(conf DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host, conf.Port, conf.Username, conf.DbName, conf.Password, conf.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ReverseAction(tx *sql.Tx,db *sql.DB, args []string, usrId int) error {
	var emptyCheck bool
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 AND first_action_args=$2", ActionsTable)
	row := db.QueryRow(query, usrId, pq.Array(args))
	err := row.Scan(&emptyCheck)
	if err == sql.ErrNoRows {
		query = fmt.Sprintf("INSERT INTO %s (second_action_id, second_action_args, second_action_type) VALUES ($1,$2,$3)", ActionsTable)
		var zeroSlice []string
		_, err = db.Exec(query, 0, pq.Array(zeroSlice), "")
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	query = fmt.Sprintf("UPDATE %s SET first_action_id=$1, first_action_args=$2, first_action_type=$3 WHERE user_id=$4", ActionsTable)
	var zeroSlice []string
	_, err = db.Exec(query, 0, pq.Array(zeroSlice), "", usrId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func InsertAction(tx *sql.Tx, name string, args []string, usrId int) error {
	actionInsertQuery := fmt.Sprintf("INSERT INTO %[1]v (user_id, first_action_type, first_action_args) VALUES ($1,$2,$3) ON CONFLICT (user_id) DO UPDATE SET"+
		" second_action_type=%[1]v.first_action_type, second_action_args=%[1]v.first_action_args, second_action_id=%[1]v.first_action_id,"+
		" first_action_type=EXCLUDED.first_action_type, first_action_args=EXCLUDED.first_action_args, first_action_id=nextval(%s)", ActionsTable, "'actions_id_seq'")
	_, err := tx.Exec(actionInsertQuery, usrId, name, pq.Array(args))
	if err != nil {
		return err
	}
	return nil
}