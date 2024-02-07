package main

import (
	"order-api/database"
	"order-api/routers"
)

func main() {
	PORT := ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}
