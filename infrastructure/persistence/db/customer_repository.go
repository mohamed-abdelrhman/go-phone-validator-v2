package db

import (
	"database/sql"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/entity"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/clients"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/utils/errors"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/validations"
	"log"
)
const(
	queryGetAllCustomer=" SELECT * from customer ;"
	queryGetCustomersByCountryCode=" SELECT * from customer where phone LIKE ? ;"
)


type CustomerRepositoryInterface interface {
	GetCustomers(filterCustomers entity.FilterCustomer) ([]entity.Customer, *errors.RestErr)
}

type customerRepository struct {
}

func NewCustomerRepository() CustomerRepositoryInterface {
	return &customerRepository{}
}

func (r *customerRepository)GetCustomers(filterCustomer entity.FilterCustomer) ([]entity.Customer, *errors.RestErr){
	var stmt *sql.Stmt
	var err error
	if filterCustomer.CountryCode=="" {
		stmt, err =clients.GetSQLClient().Prepare(queryGetAllCustomer)
	}else {
		stmt, err =clients.GetSQLClient().Prepare(queryGetCustomersByCountryCode)
		log.Println()
	}
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError("Database parsing error")
	}
	defer stmt.Close()


	var rows *sql.Rows
	if filterCustomer.CountryCode=="" {
		rows,err=stmt.Query()
	}else {
		rows,err=stmt.Query(filterCustomer.CountryCode+"%")
	}

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
		switch filterCustomer.Status {
		case "":
			results=append(results,customer)
		case "valid":
			if customer.Status {results=append(results,customer)}
		default :
			if !customer.Status {results=append(results,customer)}
		}
	}
	if len(results)==0 {
		return  nil, errors.NewNotFoundError("No Customers Found")
	}
	return results,nil
}