package main

import (
	"github.com/diaszakir/URL-Shortener-API-Go/internals/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/diaszakir/URL-Shortener-API-Go/docs"
)

// @title URL Shortener API
// @Description URL Shortener REST API project
// @Version 1.0
// @Host localhost:8080
// @BasePath /
func main() {
	s := gin.Default()

	routes.RegisterRoutes(s)

	s.Run(":8080")
}
