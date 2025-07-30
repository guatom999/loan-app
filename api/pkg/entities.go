package pkg

type (
	Statement struct {
		Id         int64  `db:"id"`
		Fullname   string `db:"fullname"`
		Salary     int    `db:"salary"`
		LoanAmount int    `db:"loan_amount"`
	}
)
