package pkg

type (
	UserCreateStateMentRequest struct {
		FullName   string `json:"fullname"`
		Salary     int    `json:"salary"`
		LoanAmount int    `json:"loan_amount"`
	}
)
