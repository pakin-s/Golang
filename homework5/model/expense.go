package model

type ExpenseTracker struct {
	ID      int     `json:"id"`
	User    string  `json:"user"`
	Income  float64 `json:"income"`
	Outcome float64 `json:"outcome"`
	Saving  float64 `json:"saving"`
}

var MockData1 = []ExpenseTracker{
	{	ID: 1,
		User:    "Pakin",
		Income:  300000.00,
		Outcome: 100000.00,
		Saving:  18000000.00},
	{	ID: 2,
		User:    "Pond",
		Income:  300000.00,
		Outcome: 100000.00,
		Saving:  18000000.00,
	},
	{	ID: 3,
		User:    "Than",
		Income:  300000.00,
		Outcome: 100000.00,
		Saving:  18000000.00,
	},
}
