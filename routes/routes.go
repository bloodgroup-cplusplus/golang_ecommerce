package routes

import (
	"github.com/bloodgroup-cplusplus/golang_ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup",controllers.Signup())
	incomingRoutes.POST("/users/login",controllers.Login())
	incomingRoutes.POST("/admin/addproduct",controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview",controllers.SearchProduct())
	incomingRoutes.GET("/users/search",controllers.SearchProductByQuery())

}
