package WishlistController

import (
	"net/http"
	dbInit "nivasProductBackendApp/DB"
	WishlistModel "nivasProductBackendApp/Model/Wishlist"
	WishlistService "nivasProductBackendApp/Service/Wishlist"

	"github.com/gin-gonic/gin"
)

func ListAllWishListProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req struct {
			UserId int `json:"userId" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
			})
			return
		}

		db, _ := dbInit.InitDB() // Get GORM DB instance
		response := WishlistService.ListAllWishListProducts(db, req.UserId)
		c.JSON(http.StatusOK, response)
	}
}

func AddProductToWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req WishlistModel.AddWishlistRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
			})
			return
		}

		db, _ := dbInit.InitDB()
		response := WishlistService.AddProductToWishlist(db, req.UserId, req.ProductId)
		c.JSON(http.StatusOK, response)
	}
}

func RemoveProductFromWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req WishlistModel.RemoveWishlistRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
			})
			return
		}

		db, _ := dbInit.InitDB()
		response := WishlistService.RemoveProductFromWishlist(db, req.WishlistId)
		c.JSON(http.StatusOK, response)
	}
}

func UpdateNotifyStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RefWishlistId int `json:"wishlistId" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
			return
		}

		db, _ := dbInit.InitDB()
		response := WishlistService.UpdateNotifyStatus(db, req.RefWishlistId)
		c.JSON(http.StatusOK, response)
	}
}

func GetProductSizesAndStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
	ProductId int `json:"productId" binding:"required"`
}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
			return
		}

		db, _ := dbInit.InitDB()
		response := WishlistService.GetProductSizesAndStock(db, req.ProductId)
		c.JSON(http.StatusOK, response)
	}
}