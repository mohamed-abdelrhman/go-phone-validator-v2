package api

import (
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/application"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/service"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/persistence/db"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/interfaces/http"
)

func CustomerUrlMapping()  {
	customers := http.NewCustomers(application.NewCustomerApp(service.NewCustomerService(db.NewCustomerRepository())))
	Router.GET("/customers", customers.GetCustomers)
}
