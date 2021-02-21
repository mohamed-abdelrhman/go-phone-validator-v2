package application

import (
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/entity"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/service"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/utils/errors"
)

type customerApp struct {
	cs service.CustomerServiceInterface
}


var _ CustomerAppInterface = &customerApp{}

type CustomerAppInterface interface {
	GetCustomers(filterCustomers entity.FilterCustomer) ([]entity.Customer, *errors.RestErr)
}
func NewCustomerApp(cs service.CustomerServiceInterface ) CustomerAppInterface {
	return &customerApp{
		cs: cs,
	}
}
func (c *customerApp) GetCustomers(filterCustomers entity.FilterCustomer) ([]entity.Customer, *errors.RestErr) {
	return c.cs.GetCustomers(filterCustomers)
}