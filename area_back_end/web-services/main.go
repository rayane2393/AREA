package main

import (
	"API/DB_management"
	"API/routes"
	"API/AREA_management"
)

func main() {
	DB_management.ConnectToDb()
	AREA_management.CheckAreaStart()
	routes.Routes()
}
