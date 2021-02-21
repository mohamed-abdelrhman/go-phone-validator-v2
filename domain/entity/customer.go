package entity

type Customer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Status    bool
}

type Customers []Customer

