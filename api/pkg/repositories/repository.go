package repositories

import "github.com/finance-app/pkg"

type (
	RepositoryInterface interface {
		CreateTransaction(stateMent *pkg.Statement) error
	}
)
