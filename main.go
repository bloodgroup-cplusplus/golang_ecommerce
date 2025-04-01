package main

import (
	"github.com/bloodgroup-cplusplus/golang_ecommerce/controllers"
	"github.com/bloodgroup-cplusplus/golang_ecommerce/database"
	"github.com/bloodgroup-cplusplus/golang_ecommerce/middleware"
	"github.com/bloodgroup-cplusplus/golang_ecommerce/routes"
	"github.com/gin-gonic/gin"
)


func main() {
	port : = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client,"Products"),database.UserData(database.Client,"Users"))
	
}