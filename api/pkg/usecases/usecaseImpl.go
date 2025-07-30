package usecases

import (
	"errors"
	"log"
	"sync"

	"github.com/finance-app/pkg"
	"github.com/finance-app/pkg/repositories"
)

type (
	queueRequest struct {
		*pkg.Statement
	}

	usecase struct {
		repo     repositories.RepositoryInterface
		queue    chan *queueRequest
		queueRes chan error
		once     sync.Once
	}
)

func NewUseCase(repo repositories.RepositoryInterface) UsecaseInterface {
	svc := &usecase{
		repo:     repo,
		queue:    make(chan *queueRequest, 100000),
		queueRes: make(chan error, 100000),
	}

	svc.once.Do(func() {
		for i := 0; i < 10; i++ {
			go func(workerId int) {
				svc.worker(workerId)
			}(i)
		}
	})
	return svc
}

func (u *usecase) worker(id int) {
	log.Printf("Worker number %d is starting", id)

	for statement := range u.queue {
		if statement == nil {
			log.Printf("Worker %d received nil job, skipping", id)
			continue
		}

		if err := u.repo.CreateTransaction(&pkg.Statement{
			Fullname:   statement.Fullname,
			Salary:     statement.Salary,
			LoanAmount: statement.LoanAmount,
		}); err != nil {
			log.Printf("Worker %d job error: %v", id, err)
			u.queueRes <- err
		} else {
			// log.Printf("Worker %d job completed successfully", id)
			u.queueRes <- nil
		}
	}
}

func (u *usecase) CheckStateMent(req *pkg.UserCreateStateMentRequest) error {

	if req.LoanAmount <= req.Salary || req.LoanAmount > req.Salary*10 {
		log.Printf("case 1 %d, %d ", req.LoanAmount, req.Salary)
		return errors.New("not approve")
	}

	if !(req.LoanAmount >= 100000 && req.LoanAmount <= 600000) {
		log.Printf("case 1 %d, %d", req.LoanAmount, req.Salary)
		return errors.New("not approve loan amount too many")
	}

	q := &queueRequest{
		Statement: &pkg.Statement{
			Fullname:   req.FullName,
			Salary:     req.Salary,
			LoanAmount: req.LoanAmount,
		},
	}

	u.queue <- q

	res := <-u.queueRes
	if res != nil {
		return errors.New(res.Error())
	}

	return nil

}
