package getProductsController

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	dbInit "nivasProductBackendApp/DB"
	getProductsModel "nivasProductBackendApp/Model/Product"
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

func ListProductByGenderOrCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getProductsModel.ListProductByGenderOrCategoryReq

		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println("Bind Error:", err) // 🧠 Add this for debugging
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Invalid request payload",
			})
			return
		}

		if req.RefCategoryGenderId == nil && req.RefParentCategoryId == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "Either refCategoryGenderId or refParentCategoryId is required",
			})
			return
		}

		db, _ := dbInit.InitDB()
		response := getProductsService.ListProductByGenderOrCategoryService(db, req)
		c.JSON(http.StatusOK, response)
	}
}

func ListCategoryPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := dbInit.InitDB() // Get GORM DB instance
		response := getProductsService.ListCategoryPageService(db)
		c.JSON(http.StatusOK, response)
	}
}

func GetParentCategory() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get ID from URL param
        RefCategoryGenderIdStr := c.Param("id")
        RefCategoryGenderId, err := strconv.Atoi(RefCategoryGenderIdStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status":  false,
                "message": "Invalid ID",
            })
            return
        }

        db, _ := dbInit.InitDB()
        response := getProductsService.GetParentCategoryService(db, RefCategoryGenderId)

        c.JSON(http.StatusOK, response)
    }
}
func GetSubCategoryByParentCategory() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get ID from URL param
         RefSubCategoryIdStr := c.Param("id")
        RefSubCategoryId, err := strconv.Atoi(RefSubCategoryIdStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status":  false,
                "message": "Invalid ID",
            })
            return
        }

        db, _ := dbInit.InitDB()
        response := getProductsService.GetSubCategoryByParentCategoryService(db, RefSubCategoryId)

        c.JSON(http.StatusOK, response)
    }
}

func GetProductsByParentCategory() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get ID from URL param
        RefParentCategoryIdStr := c.Param("id")
        RefParentCategoryId, err := strconv.Atoi(RefParentCategoryIdStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status":  false,
                "message": "Invalid ID",
            })
            return
        }

        db, _ := dbInit.InitDB()
        response := getProductsService.GetProductsByParentCategoryService(db, RefParentCategoryId)

        c.JSON(http.StatusOK, response)
    }
}

func GetProductsBySubCategory() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get ID from URL param
        RefSUbCategoryIdStr := c.Param("id")
        RefSubCategoryId, err := strconv.Atoi(RefSUbCategoryIdStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status":  false,
                "message": "Invalid ID",
            })
            return
        }

        db, _ := dbInit.InitDB()
        response := getProductsService.GetProductsBySubCategoryService(db, RefSubCategoryId)

        c.JSON(http.StatusOK, response)
    }
}

func ListNewArrivals() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, _ := dbInit.InitDB() // Get GORM DB instance
		response := getProductsService.ListNewArrivalsService(db)
		c.JSON(http.StatusOK, response)
	}
}

func GetNewArrivalProductsByBrand() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get ID from URL param
        RefApplicationIdStr := c.Param("id")
        RefApplicationId, err := strconv.Atoi(RefApplicationIdStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status":  false,
                "message": "Invalid ID",
            })
            return
        }

        db, _ := dbInit.InitDB()
        response := getProductsService.GetNewArrivalProductsByBrandService(db, RefApplicationId)

        c.JSON(http.StatusOK, response)
    }
}