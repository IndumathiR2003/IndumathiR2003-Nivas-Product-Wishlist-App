package getProductsRoutes

import (
	"github.com/gin-gonic/gin"

	getProductsController "nivasProductBackendApp/Controller/Product"

)

func ProductCategory(router *gin.Engine) {
	route := router.Group("/products/api/v1/list")
	route.GET("/listBrandsWithLogo", getProductsController.ListBrandsandLogo())
	route.POST("/listParentCategoryByGender", getProductsController.ListParentCategoryByGender())
	route.GET("/listParentandSubCategories", getProductsController.ListParentandSubCategories())
	route.GET("/listProducts", getProductsController.ListProducts())
    route.GET("/getProductDetails/:id", getProductsController.GetProductDetails())

}
