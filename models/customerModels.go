package models

type Customer struct {
	CustomerId   int    `json:"customer_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	CustomerType string `json:"customer_type"`
	JoinDate     string `json:"join_date"`
}

type CustomersListResponse struct {
	Data       []Customer `json:"customers"`
	Pagination Pagination `json:"pagination"`
}
