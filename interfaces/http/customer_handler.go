package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/application"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/entity"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/utils/errors"
	"log"
	"net/http"
)

type CustomerHandlerInterface interface {
	GetCustomers(c *gin.Context)
}

type customerHandler struct {
	ca application.CustomerAppInterface
}

func NewCustomers(ca application.CustomerAppInterface ) CustomerHandlerInterface {
	return &customerHandler{
		ca: ca,
	}
}

func (ch *customerHandler) GetCustomers(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	c.Next()
	var filterCustomers entity.FilterCustomer
	filterCustomers.CountryCode = c.Query("country_code")
	filterCustomers.Status = c.Query("status")
	log.Println(filterCustomers)
	customers := entity.Customers{}
	var err *errors.RestErr
	customers, err = ch.ca.GetCustomers(filterCustomers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, customers)
}

