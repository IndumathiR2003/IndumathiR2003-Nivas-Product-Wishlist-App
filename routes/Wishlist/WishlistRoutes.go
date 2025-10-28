package WishlistRoutes

import (
	WishlistController "nivasProductBackendApp/Controller/Wishlist"

	"github.com/gin-gonic/gin"
)

func WishlistRoutes(router *gin.Engine) {
	route := router.Group("/wishlist/api/v1/")
	route.POST("/listAllWishListProducts", WishlistController.ListAllWishListProducts())
	route.POST("/addProductToWishlist", WishlistController.AddProductToWishlist())
	route.POST("/removeProductFromWishlist", WishlistController.RemoveProductFromWishlist())
	route.POST("/updateWishlistNotifyStatus", WishlistController.UpdateNotifyStatus())
	route.POST("/getProductSizesAndStock", WishlistController.GetProductSizesAndStock())
}