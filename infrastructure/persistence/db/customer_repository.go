package db

import (
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/entity"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/clients"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/utils/errors"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/validations"
	"log"
)
const(
	queryGetAllCustomer=" SELECT * from customer ;"
)


type CustomerRepositoryInterface interface {
	GetCustomers() ([]entity.Customer, *errors.RestErr)
}

type customerRepository struct {
}

func NewCustomerRepository() CustomerRepositoryInterface {
	return &customerRepository{}
}

func (r *customerRepository)GetCustomers() ([]entity.Customer, *errors.RestErr){
	stmt, err :=clients.GetSQLClient().Prepare(queryGetAllCustomer)
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()
	rows,err:=stmt.Query()
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError("Database parsing error")
	}
	defer rows.Close()
	results:=make([]entity.Customer,0)
	for rows.Next(){
		var customer entity.Customer

		if err :=rows.Scan(&customer.ID,&customer.Name,&customer.Phone);err!=nil{
			log.Println(err)
			return nil, errors.NewInternalServerError("Database parsing error")
		}
		customer.Status =validations.ValidatePhone(customer.Phone)
		results=append(results,customer)
	}
	if len(results)==0 {
		return  nil, errors.NewNotFoundError("No Customers Found")
	}
	return results,nil
}