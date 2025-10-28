package getProductsService

import (
	"fmt"

	"gorm.io/gorm"

	logger "nivasProductBackendApp/Helper/Logger"
	minioService "nivasProductBackendApp/Helper/MinIo"
	getProductsModel "nivasProductBackendApp/Model/Product"
	getProductsQuery "nivasProductBackendApp/Query/Product"
)

func ListBrandsandLogoService(db *gorm.DB) getProductsModel.ListBrandsWithLogoOptionsResp {
	fmt.Println("Registered routes:")

	log := logger.InitLogger()
	var brandList []getProductsModel.BrandOptionsQueryResp

	err := db.Raw(getProductsQuery.GetBrandListQuery).Scan(&brandList).Error
	if err != nil {
		log.Error("Error fetching brand list: " + err.Error())
		return getProductsModel.ListBrandsWithLogoOptionsResp{
			Status:  false,
			Message: "Failed to get brand list",
		}
	}

	// Generate MinIO URLs
	for i := range brandList {
		if brandList[i].RefLogoFileName != "" {
			url, err := minioService.GetFileURL(brandList[i].RefLogoFileName, 60)
			if err == nil {
				brandList[i].RefLogoFileName = url
			}
		}
		if brandList[i].RefCelebrityFileName != "" {
			url, err := minioService.GetFileURL(brandList[i].RefCelebrityFileName, 60)
			if err == nil {
				brandList[i].RefCelebrityFileName = url
			}
		}
	}

	return getProductsModel.ListBrandsWithLogoOptionsResp{
		Status:       true,
		Message:      "Brand list fetched successfully",
		BrandOptions: brandList,
	}
}

func ListParentCategorybyGenderService(db *gorm.DB, genderId int) getProductsModel.ListParentCategoryResp {
	fmt.Println("Registered routes:", genderId)
	log := logger.InitLogger()
	var categoryList []getProductsModel.ParentCategoryResp

	err := db.Raw(getProductsQuery.GetParentCategorybyGenderQuery, genderId).Scan(&categoryList).Error
	if err != nil {
		log.Error("Error fetching category list: " + err.Error())
		return getProductsModel.ListParentCategoryResp{
			Status:  false,
			Message: "Failed to get parent categories",
		}
	}

	// Generate MinIO URLs for images if available
	for i := range categoryList {
		if categoryList[i].RefParentCategoryImagePath != "" {
			url, err := minioService.GetFileURL(categoryList[i].RefParentCategoryImagePath, 60)
			if err == nil {
				categoryList[i].RefParentCategoryImagePath = url
			}
		}
	}

	return getProductsModel.ListParentCategoryResp{
		Status:           true,
		Message:          "Parent categories fetched successfully",
		ParentCategories: categoryList,
	}
}

func ListParentandSubCategoriesService(db *gorm.DB) getProductsModel.ListParentCategoryWithSubResp {
	log := logger.InitLogger()

	// Step 1: Fetch all parent categories (raw)
	var rawParents []getProductsModel.ParentCategoryRaw
	parentQuery := `
        SELECT
            pcm."refParentCateroryMapId",
            g."refCategoryGenderId",
            g."refCategoryGenderName" AS "genderName",
            pc."refParentCategoryId",
            pc."refParentCategoryName" AS "parentCategoryName",
            pc."refParentCategoryImagePath",
            pcm."refIfVisible" AS parentVisible
        FROM productcategory."refParentCategoryMapping" pcm
        JOIN productcategory."refCategoryGender" g
            ON pcm."refCategoryGenderId" = g."refCategoryGenderId"
        JOIN productcategory."refParentCategory" pc
            ON pcm."refParentCategoryId" = pc."refParentCategoryId"
        WHERE pcm."refIfVisible" = TRUE
        ORDER BY g."refCategoryGenderName", pc."refParentCategoryName";
    `
	if err := db.Raw(parentQuery).Scan(&rawParents).Error; err != nil {
		log.Error("Error fetching parent categories: " + err.Error())
		return getProductsModel.ListParentCategoryWithSubResp{
			Status:  false,
			Message: "Failed to fetch parent categories",
		}
	}

	// Step 2: Populate subcategories for each parent
	var parentsWithSubs []getProductsModel.ParentCategoryWithSub
	for _, p := range rawParents {
		var subs []getProductsModel.SubCategory
		subQuery := `
            SELECT
                sc."refSubCategoryId",
                sc."refSubCategory" AS "subCategoryName"
            FROM productcategory."refParentSubCategoryMapping" pscm
            JOIN productcategory."refSubCategory" sc
                ON pscm."refSubCategoryId" = sc."refSubCategoryId"
            WHERE pscm."refParentGenderMappingId" = ?
              AND sc."refIfVisible" = TRUE;
        `
		if err := db.Raw(subQuery, p.RefParentCateroryMapId).Scan(&subs).Error; err != nil {
			log.Error("Error fetching subcategories for parent ID " +
				string(p.RefParentCateroryMapId) + ": " + err.Error())
			continue
		}

		parentsWithSubs = append(parentsWithSubs, getProductsModel.ParentCategoryWithSub{
			RefParentCateroryMapId:     p.RefParentCateroryMapId,
			RefCategoryGenderId:        p.RefCategoryGenderId,
			GenderName:                 p.GenderName,
			ParentCategoryId:           p.ParentCategoryId,
			ParentCategoryName:         p.ParentCategoryName,
			RefParentCategoryImagePath: p.RefParentCategoryImagePath,
			ParentVisible:              p.ParentVisible,
			SubCategories:              subs,
		})
	}

	return getProductsModel.ListParentCategoryWithSubResp{
		Status:           true,
		Message:          "Parent categories fetched successfully",
		ParentCategories: parentsWithSubs,
	}
}

func ListProductsService(db *gorm.DB) getProductsModel.ProductDefaultCatalogResponseOptionsResp {
	fmt.Println("Registered routes:")
	log := logger.InitLogger()
	var productList []getProductsModel.ProductDefaultCatalogResponse
	fmt.Println("Registered routes:", productList)
	err := db.Raw(getProductsQuery.GetVisibleProductCatalogQuery).Scan(&productList).Error
	if err != nil {
		log.Error("Error fetching Product: " + err.Error())
		return getProductsModel.ProductDefaultCatalogResponseOptionsResp{
			Status:  false,
			Message: "Failed to get Product",
		}
	}

	// Generate MinIO URLs
	for i := range productList {
		if productList[i].RefProductDetailAngleImage != "" {
			url, err := minioService.GetFileURL(productList[i].RefProductDetailAngleImage, 60)
			if err == nil {
				productList[i].RefProductDetailAngleImage = url
			}
		}

	}

	return getProductsModel.ProductDefaultCatalogResponseOptionsResp{
		Status:         true,
		Message:        "Product fetched successfully",
		ProductOptions: productList,
	}
}

func GetProductDetailsService(db *gorm.DB, RefProductId int) getProductsModel.ListProductsOptions {
	fmt.Println("Registered routes:", RefProductId)
	log := logger.InitLogger()
	var ProductList []getProductsModel.ProductsResp

	err := db.Raw(getProductsQuery.GetproductDetailsQuery, RefProductId).Scan(&ProductList).Error
	if err != nil {
		log.Error("Error fetching category list: " + err.Error())
		return getProductsModel.ListProductsOptions{
			Status:  false,
			Message: "Failed to get parent categories",
		}
	}

	// Generate MinIO URLs for images if available
	for i := range ProductList {
		if ProductList[i].RefProductDetailAngleImage != "" {
			url, err := minioService.GetFileURL(ProductList[i].RefProductDetailAngleImage, 60)
			if err == nil {
				ProductList[i].RefProductDetailAngleImage = url
			}
		}
		if ProductList[i].RefProductFrontImage != "" {
			url, err := minioService.GetFileURL(ProductList[i].RefProductFrontImage, 60)
			if err == nil {
				ProductList[i].RefProductFrontImage = url
			}
		}
		if ProductList[i].RefProductSideImage != "" {
			url, err := minioService.GetFileURL(ProductList[i].RefProductSideImage, 60)
			if err == nil {
				ProductList[i].RefProductSideImage = url
			}
		}
		if ProductList[i].RefProductBackImage != "" {
			url, err := minioService.GetFileURL(ProductList[i].RefProductBackImage, 60)
			if err == nil {
				ProductList[i].RefProductBackImage = url
			}
		}
		if ProductList[i].RefProductAdditionalImage != "" {
			url, err := minioService.GetFileURL(ProductList[i].RefProductAdditionalImage, 60)
			if err == nil {
				ProductList[i].RefProductAdditionalImage = url
			}
		}
		if ProductList[i].RefProductVideo != "" {
			url, err := minioService.GetFileURL(ProductList[i].RefProductVideo, 60)
			if err == nil {
				ProductList[i].RefProductVideo = url
			}
		}
	}

	return getProductsModel.ListProductsOptions{
		Status:   true,
		Message:  "Products fetched successfully",
		Products: ProductList,
	}
}

func ListProductByGenderOrCategoryService(db *gorm.DB, req getProductsModel.ListProductByGenderOrCategoryReq) getProductsModel.ListProductByGenderResp {
	log := logger.InitLogger()
	var productsList []getProductsModel.ProductByGenderResp
	var err error

	if req.RefCategoryGenderId != nil {
		err = db.Raw(getProductsQuery.GetProductsByGender, *req.RefCategoryGenderId).Scan(&productsList).Error
	} else if req.RefParentCategoryId != nil {
		err = db.Raw(getProductsQuery.GetProductsByParentCategory, *req.RefParentCategoryId).Scan(&productsList).Error
	}

	if err != nil {
		log.Error("Error fetching products: " + err.Error())
		return getProductsModel.ListProductByGenderResp{
			Status:  false,
			Message: "Failed to fetch products",
		}
	}

	// Generate signed URLs
	for i := range productsList {
		if productsList[i].RefProductDetailAngleImage != "" {
			url, err := minioService.GetFileURL(productsList[i].RefProductDetailAngleImage, 60)
			if err == nil {
				productsList[i].RefProductDetailAngleImage = url
			}
		}
	}

	return getProductsModel.ListProductByGenderResp{
		Status:          true,
		Message:         "Products fetched successfully",
		ProductByGender: productsList,
	}
}

func ListCategoryPageService(db *gorm.DB) getProductsModel.ListCategoryPageOptionsResp {
	log := logger.InitLogger()
	var categoryList []getProductsModel.ListCategoryPageResponse
	fmt.Println("Registered routes:", categoryList)
	err := db.Raw(getProductsQuery.ListCategoryPageParentCategroryQuery).Scan(&categoryList).Error
	if err != nil {
		log.Error("Error fetching Product: " + err.Error())
		return getProductsModel.ListCategoryPageOptionsResp{
			Status:  false,
			Message: "Failed to get Product",
		}
	}

	// Generate MinIO URLs
	for i := range categoryList {
		if categoryList[i].RefGenderImagePath != "" {
			url, err := minioService.GetFileURL(categoryList[i].RefGenderImagePath, 60)
			if err == nil {
				categoryList[i].RefGenderImagePath = url
			}
		}
		if categoryList[i].RefCategoryImagePath != "" {
			url, err := minioService.GetFileURL(categoryList[i].RefCategoryImagePath, 60)
			if err == nil {
				categoryList[i].RefCategoryImagePath = url
			}
		}

	}

	return getProductsModel.ListCategoryPageOptionsResp{
		Status:  true,
		Message: "Product fetched successfully",
		Data:    categoryList,
	}
}

func GetParentCategoryService(db *gorm.DB, RefCategoryGenderId int) getProductsModel.GetParentCategoryOptions {
	fmt.Println("Registered routes:", RefCategoryGenderId)
	log := logger.InitLogger()
	var ParentCategoryList []getProductsModel.GetParentCategoryResp
	fmt.Println("Registered routes:", ParentCategoryList)

	err := db.Raw(getProductsQuery.ListParentCategrorywithSubcategoryQuery, RefCategoryGenderId).Scan(&ParentCategoryList).Error
	if err != nil {
		log.Error("Error fetching category list: " + err.Error())
		return getProductsModel.GetParentCategoryOptions{
			Status:  false,
			Message: "Failed to get parent categories",
		}
	}

	return getProductsModel.GetParentCategoryOptions{
		Status:             true,
		Message:            "ParentCategoryList fetched successfully",
		ParentCategoryList: ParentCategoryList,
	}
}

func GetSubCategoryByParentCategoryService(db *gorm.DB, RefSubCategoryId int) getProductsModel.GetSubCategoryOptions {
	fmt.Println("Registered routes:", RefSubCategoryId)
	log := logger.InitLogger()
	var SubCategoryList []getProductsModel.GetSubCategoryResp
	fmt.Println("Registered routes:", SubCategoryList)

	err := db.Raw(getProductsQuery.ListSubcategoryQuery, RefSubCategoryId).Scan(&SubCategoryList).Error
	if err != nil {
		log.Error("Error fetching category list: " + err.Error())
		return getProductsModel.GetSubCategoryOptions{
			Status:  false,
			Message: "Failed to get parent categories",
		}
	}

	return getProductsModel.GetSubCategoryOptions{
		Status:          true,
		Message:         "ParentCategoryList fetched successfully",
		SubCategoryList: SubCategoryList,
	}
}

func GetProductsByParentCategoryService(db *gorm.DB, RefParentCategoryId int) getProductsModel.ProductsByParentCategoryOptions {
	fmt.Println("Registered routes:", RefParentCategoryId)
	log := logger.InitLogger()
	var ProductsByParentCategory []getProductsModel.ProductsByParentCategoryResp
	fmt.Println("Registered routes:", ProductsByParentCategory)

	err := db.Raw(getProductsQuery.ListProductsByParentCategory, RefParentCategoryId).Scan(&ProductsByParentCategory).Error
	if err != nil {
		log.Error("Error fetching category list: " + err.Error())
		return getProductsModel.ProductsByParentCategoryOptions{
			Status:  false,
			Message: "Failed to get products by parent categories",
		}
	}

	for i := range ProductsByParentCategory {
		if ProductsByParentCategory[i].RefProductDetailAngleImage != "" {
			url, err := minioService.GetFileURL(ProductsByParentCategory[i].RefProductDetailAngleImage, 60)
			if err == nil {
				ProductsByParentCategory[i].RefProductDetailAngleImage = url
			}
		}

	}

	return getProductsModel.ProductsByParentCategoryOptions{
		Status:                   true,
		Message:                  "products by parent categories fetched successfully",
		ProductsByParentCategory: ProductsByParentCategory,
	}
}

func GetProductsBySubCategoryService(db *gorm.DB, RefSubCategoryId int) getProductsModel.ProductsBySubCategoryOptions {
	fmt.Println("Registered routes:", RefSubCategoryId)
	log := logger.InitLogger()
	var ProductsBySubCategory []getProductsModel.ProductsBySubCategoryResp
	fmt.Println("Registered routes:", ProductsBySubCategory)

	err := db.Raw(getProductsQuery.ListProductsByParentCategory, RefSubCategoryId).Scan(&ProductsBySubCategory).Error
	if err != nil {
		log.Error("Error fetching products list: " + err.Error())
		return getProductsModel.ProductsBySubCategoryOptions{
			Status:  false,
			Message: "Failed to get products by sub categories",
		}
	}

	for i := range ProductsBySubCategory {
		if ProductsBySubCategory[i].RefProductDetailAngleImage != "" {
			url, err := minioService.GetFileURL(ProductsBySubCategory[i].RefProductDetailAngleImage, 60)
			if err == nil {
				ProductsBySubCategory[i].RefProductDetailAngleImage = url
			}
		}

	}

	return getProductsModel.ProductsBySubCategoryOptions{
		Status:                true,
		Message:               "products by sub categories fetched successfully",
		ProductsBySubCategory: ProductsBySubCategory,
	}
}

func ListNewArrivalsService(db *gorm.DB) getProductsModel.ListNewArrivalsResponseOptionsResp {
	log := logger.InitLogger()
	var productList []getProductsModel.ListNewArrivalsResponse
	fmt.Println("productList", productList)

	err := db.Raw(getProductsQuery.ListNewArrivalsQuery).Scan(&productList).Error
	if err != nil {
		log.Error("Error fetching Product: " + err.Error())
		return getProductsModel.ListNewArrivalsResponseOptionsResp{
			Status:  false,
			Message: "Failed to get Product",
		}
	}
	fmt.Println("Registered routes:", productList[0].RefLogoFileName)
	// Generate MinIO URLs
	for i := range productList {
		if productList[i].RefProductDetailAngleImage != "" {
			url, err := minioService.GetFileURL(productList[i].RefProductDetailAngleImage, 60)
			if err == nil {
				productList[i].RefProductDetailAngleImage = url
			}
		}
		if productList[i].RefLogoFileName != "" {
			url, err := minioService.GetFileURL(productList[i].RefLogoFileName, 60)
			if err == nil {
				productList[i].RefLogoFileName = url
			}
		}

	}

	return getProductsModel.ListNewArrivalsResponseOptionsResp{
		Status:  true,
		Message: "Product fetched successfully",
		Data:    productList,
	}
}

func GetNewArrivalProductsByBrandService(db *gorm.DB, RefApplicationId int) getProductsModel.GetNewArrivalProductsByBrandOptionsResp {
	fmt.Println("Registered routes:", RefApplicationId)
	log := logger.InitLogger()
	var NewArrivalsByBrand []getProductsModel.GetNewArrivalProductsByBrandResponse
	fmt.Println("Registered routes:", NewArrivalsByBrand)

	err := db.Raw(getProductsQuery.GetNewArrivalProductsByBrandQuery, RefApplicationId).Scan(&NewArrivalsByBrand).Error
	if err != nil {
		log.Error("Error fetching products list: " + err.Error())
		return getProductsModel.GetNewArrivalProductsByBrandOptionsResp{
			Status:  false,
			Message: "Failed to get products by sub categories",
		}
	}

	for i := range NewArrivalsByBrand {
		if NewArrivalsByBrand[i].RefProductDetailAngleImage != "" {
			url, err := minioService.GetFileURL(NewArrivalsByBrand[i].RefProductDetailAngleImage, 60)
			if err == nil {
				NewArrivalsByBrand[i].RefProductDetailAngleImage = url
			}
		}

	}

	return getProductsModel.GetNewArrivalProductsByBrandOptionsResp{
		Status:                       true,
		Message:                      "products by sub categories fetched successfully",
		GetNewArrivalProductsByBrand: NewArrivalsByBrand,
	}
}
