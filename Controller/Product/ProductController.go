package getProductsController

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dbInit "nivasProductBackendApp/DB"
	getProductsService "nivasProductBackendApp/Service/Product"
)

func ListBrandsandLogo() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := dbInit.InitDB() // Get GORM DB instance
		response := getProductsService.ListBrandsandLogoService(db)
		c.JSON(http.StatusOK, response)
	}
}

func ListParentCategoryByGender() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			RefCategoryGenderId int `json:"refCategoryGenderId" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
			})
			return
		}

		db, _ := dbInit.InitDB()
		response := getProductsService.ListParentCategorybyGenderService(db, req.RefCategoryGenderId)
		c.JSON(http.StatusOK, response)
	}
}

func ListParentandSubCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := dbInit.InitDB() // Get GORM DB instance
		response := getProductsService.ListParentandSubCategoriesService(db)
		c.JSON(http.StatusOK, response)
	}
}

func ListProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := dbInit.InitDB() // Get GORM DB instance
		response := getProductsService.ListProductsService(db)
		c.JSON(http.StatusOK, response)
	}
}

func GetProductDetails() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get product ID from URL param
		refProductIdStr := c.Param("id")
		refProductId, err := strconv.Atoi(refProductIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid product ID",
			})
			return
		}

		db, _ := dbInit.InitDB()
		response := getProductsService.GetProductDetailsService(db, refProductId)
		c.JSON(http.StatusOK, response)
	}
}
