package CartController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	dbInit "nivasProductBackendApp/DB"

	logger "nivasProductBackendApp/Helper/Logger"
	CartModel "nivasProductBackendApp/Model/Cart"
	CartService "nivasProductBackendApp/Service/Cart"
)

// 📦 Add product to cart
func AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logger.InitLogger()

		var payload CartModel.AddToCartRequest

		if err := c.ShouldBindJSON(&payload); err != nil {
			log.Errorf("Invalid request payload: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
				"error":   err.Error(),
			})
			return
		}

		db, _ := dbInit.InitDB()
		result := CartService.AddToCart(db, payload)
		c.JSON(http.StatusOK, result)
	}
}

func ListUserCart() gin.HandlerFunc {
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

		db, _ := dbInit.InitDB()
		response := CartService.ListUserCart(db, req.UserId)
		c.JSON(http.StatusOK, response)
	}
}

func UpdateUserCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RefProductId  int  `json:"refProductId" binding:"required"`
			RefCartId   int `json:"refCartId" binding:"required"`
			RefQuantity int `json:"refQuantity" binding:"required"`
			RefUpdateBy int `json:"refUserId" binding:"required"`
		}

		// Validate JSON input
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
			})
			return
		}

		// Initialize DB
		db, _ := dbInit.InitDB()

		// Call service function
		response := CartService.UpdateCart(db, CartModel.UpdateCartRequest{
			RefProductId:  req.RefProductId,
			RefCartId:   req.RefCartId,
			RefQuantity: req.RefQuantity,
			RefUpdateBy: req.RefUpdateBy,
		})

		// Send response
		c.JSON(http.StatusOK, response)
	}
}
