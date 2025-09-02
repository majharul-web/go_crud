package main

import (
	"go_crud/database"
	"go_crud/routes"
)

func main() {
    database.ConnectDB()
    r := routes.SetupRouter()
    r.Run(":8080") // start server on localhost:8080
}
