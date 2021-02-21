package api

import (
	cors "github.com/rs/cors/wrapper/gin"
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
	Router.Use(cors.AllowAll())
	log.Fatal(Router.Run(":8000"))
}
