package service

import (
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/entity"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/persistence/db"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/utils/errors"
)

type CustomerServiceInterface interface {
	GetCustomers() ([]entity.Customer, *errors.RestErr)
}


type customerService struct {
	ur db.CustomerRepositoryInterface
}
func NewCustomerService(ur db.CustomerRepositoryInterface ) CustomerServiceInterface {
	return &customerService{
		ur: ur,
	}

}
func (s *customerService)GetCustomers() ([]entity.Customer, *errors.RestErr){
	return s.ur.GetCustomers()
}