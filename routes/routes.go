package routes

import (
	 "github.com/boydeptrai/ecommerce-go-project/controllers"
	 "github.com/gin-gonic/gin"
)

func UserRouter(incomingRoutes *gin.Engine) {
   incomingRoutes.POST("/users/signup", controllers.SignUp())
   incomingRoutes.POST("/users/login", controllers.Login())
   incomingRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin)
   incomingRoutes.GET("/users/productview", controllers.SearchProduct)
   incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
