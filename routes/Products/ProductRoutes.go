package getProductsRoutes

import (
	"github.com/gin-gonic/gin"

	getProductsController "nivasProductBackendApp/Controller/Product"

)

func ProductCategory(router *gin.Engine) {
	route := router.Group("/app-product-wishlist-cart/api/v1/list")
	route.GET("/listBrandsWithLogo", getProductsController.ListBrandsandLogo())
	route.POST("/listParentCategoryByGender", getProductsController.ListParentCategoryByGender())
	route.GET("/listParentandSubCategories", getProductsController.ListParentandSubCategories())
	route.GET("/listProducts", getProductsController.ListProducts())
	//pending
	route.GET("/getProductDetails/:id", getProductsController.GetProductDetails())
	route.POST("/listProductByGender", getProductsController.ListProductByGenderOrCategory())



	// completed
	route.GET("/listCategoryPage", getProductsController.ListCategoryPage())
	route.GET("/getParentByGenderCategory/:id", getProductsController.GetParentCategory())
	route.GET("/getSubCategoryByParentCategory/:id", getProductsController.GetSubCategoryByParentCategory())
	route.GET("/getProductsByParentCategory/:id", getProductsController.GetProductsByParentCategory())
	route.GET("/getProductsBySubCategory/:id", getProductsController.GetProductsBySubCategory())
	route.GET("/listNewArrivals", getProductsController.ListNewArrivals())
	route.GET("/getNewArrivalProductsByBrand/:id", getProductsController.GetNewArrivalProductsByBrand())
}
