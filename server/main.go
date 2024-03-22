package main

import (
	"blogs/server/db"
	"blogs/server/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" //PostgresSQL driver
)

func main() {
	database, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	defer database.Close()

	handler.SetupRoutes(r, database) // Assuming you have a function named SetupRoutes in your handler package

	r.Run(":8080")
}
