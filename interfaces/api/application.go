package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-abdelrhman/go-phone-validator-v2/infrastructure/clients"
	"log"
)

var(
	Router=gin.Default()
)

func StartApplication() {
	clients.GetSQLClient()
	CustomerUrlMapping()
	log.Fatal(Router.Run(":8000"))
}
