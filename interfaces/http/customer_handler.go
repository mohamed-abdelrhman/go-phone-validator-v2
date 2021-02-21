package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/application"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/domain/entity"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/utils/errors"
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
	customers := entity.Customers{}
	var err *errors.RestErr
	customers, err = ch.ca.GetCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, customers)
}
