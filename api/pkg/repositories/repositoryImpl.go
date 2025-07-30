package repositories

import (
	"github.com/finance-app/pkg"
	"github.com/jmoiron/sqlx"
)

type (
	repository struct {
		db *sqlx.DB
	}
)

func NewRepository(db *sqlx.DB) RepositoryInterface {
	return &repository{db: db}
}

func (r *repository) CreateTransaction(statement *pkg.Statement) error {

	querystring := `INSERT INTO user_statement (fullname, salary, loan_amount) VALUEs ($1 , $2 , $3)`

	_, err := r.db.Exec(querystring, statement.Fullname, statement.Salary, statement.LoanAmount)
	if err != nil {
		return err
	}

	return nil
}
