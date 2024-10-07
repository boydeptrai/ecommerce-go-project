package main

import (
	"log"
	"os"

	"github.com/boydeptrai/ecommerce-go-project/controllers"
	"github.com/boydeptrai/ecommerce-go-project/database"
	"github.com/boydeptrai/ecommerce-go-project/middleware"
	"github.com/boydeptrai/ecommerce-go-project/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	port := os.Getenv("PORT")
	if port ==""{
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client,"Products"), database.UserData(database.Client))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouter(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout",app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}