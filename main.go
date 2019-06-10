package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/melardev/GoGormApiCrudPagination/infrastructure"
	"github.com/melardev/GoGormApiCrudPagination/routes"
	"github.com/melardev/GoGormApiCrudPagination/seeds"
	"net/http"
	"os"
)

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}

	database := infrastructure.OpenDbConnection()
	defer database.Close()
	seeds.Seed(database)

	routes.RegisterRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		println(err.Error())
	}
}
