package main

import (
	"example.com/RestAPI/db"
	"example.com/RestAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	defer db.DB.Close()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost port 8080

}
