package CartRoutes

import (
	CartController "nivasProductBackendApp/Controller/Cart"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine) {
	route := router.Group("/cart/api/v1/")
	route.POST("/addToCart", CartController.AddToCart())
	route.POST("/listUserCart", CartController.ListUserCart())
	route.POST("/updateUserCart", CartController.UpdateUserCart())
}
