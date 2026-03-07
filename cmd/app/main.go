package main

import (
	"github.com/diaszakir/URL-Shortener-API-Go/internals/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	routes.RegisterRoutes(s)

	s.Run(":8080")
}
