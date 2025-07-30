package usecases

import "github.com/finance-app/pkg"

type (
	UsecaseInterface interface {
		CheckStateMent(req *pkg.UserCreateStateMentRequest) error
	}
)
