package entity

type Customer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Status    bool `json:"status"`
}

type Customers []Customer

type FilterCustomer struct {
	CountryCode string
	Status string
}