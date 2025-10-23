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
		Status:           true,
		Message:          "Products fetched successfully",
		Products: ProductList,
	}
}