package WishlistService

import (
	"fmt"
	logger "nivasProductBackendApp/Helper/Logger"
	WishlistModel "nivasProductBackendApp/Model/Wishlist"
	WishlistQuery "nivasProductBackendApp/Query/Wishlist"
	minioService "nivasProductBackendApp/Helper/MinIo"
	"gorm.io/gorm"
)

func ListAllWishListProducts(db *gorm.DB, userId int) WishlistModel.ListWishlistModelResponse {
	fmt.Println("Userrr Id:", userId)
	log := logger.InitLogger()

	var wishList []WishlistModel.ListWishlistModel

	err := db.Raw(WishlistQuery.GetWishlistProductsQuery, userId).Scan(&wishList).Error
	if err != nil {
		log.Error("DB Error:", err)
		return WishlistModel.ListWishlistModelResponse{
			Status:   false,
			Message:  "Failed to fetch wishlist",
			Products: []WishlistModel.ListWishlistModel{},
		}
	}

	// 🔹 Add hardcoded fields
	for i := range wishList {
		wishList[i].RefRatingStar = 4.5
		wishList[i].RefRatingTotalReviewers = 200 // or any number you want
		if wishList[i].RefProductImage != "" {
            url, err := minioService.GetFileURL(wishList[i].RefProductImage, 60)
            if err == nil {
                wishList[i].RefProductImage = url
            }
        }
	}

	return WishlistModel.ListWishlistModelResponse{
		Status:   true,
		Message:  "Wishlist fetched successfully",
		Products: wishList,
	}
}

func AddProductToWishlist(db *gorm.DB, userId int, productId int) WishlistModel.AddWishlistResponse {
	log := logger.InitLogger()
	var wishlistId int

	err := db.Raw(WishlistQuery.AddToWishlistQuery, userId, productId, userId, "2025").Scan(&wishlistId).Error
	if err != nil {
		log.Errorf("Error adding product to wishlist: %v", err)
		return WishlistModel.AddWishlistResponse{
			Status:  false,
			Message: "Failed to add product to wishlist",
		}
	}

	return WishlistModel.AddWishlistResponse{
		Status:  true,
		Message: "Product added to wishlist successfully",
	}
}

func RemoveProductFromWishlist(db *gorm.DB, wishlistId int) WishlistModel.RemoveWishlistResponse {
	log := logger.InitLogger()
	var removedId int

	err := db.Raw(WishlistQuery.RemoveFromWishlistQuery, wishlistId).Scan(&removedId).Error
	if err != nil {
		log.Errorf("Error removing wishlist item: %v", err)
		return WishlistModel.RemoveWishlistResponse{
			Status:  false,
			Message: "Failed to remove wishlist item",
		}
	}

	if removedId == 0 {
		return WishlistModel.RemoveWishlistResponse{
			Status:  false,
			Message: "Wishlist item not found",
		}
	}

	return WishlistModel.RemoveWishlistResponse{
		Status:  true,
		Message: "Product removed from wishlist successfully",
	}
}

func UpdateNotifyStatus(db *gorm.DB, refWishlistId int) WishlistModel.WishlistUpdateResponse {
	log := logger.InitLogger()

	result := db.Exec(WishlistQuery.UpdateNotifyStatusQuery, refWishlistId)

	if result.Error != nil {
		log.Errorf("Failed to update notify status: %v", result.Error)
		return WishlistModel.WishlistUpdateResponse{
			Status:  false,
			Message: "Failed to update notify status",
		}
	}

	if result.RowsAffected == 0 {
		return WishlistModel.WishlistUpdateResponse{
			Status:  false,
			Message: "Wishlist item not found",
		}
	}

	return WishlistModel.WishlistUpdateResponse{
		Status:  true,
		Message: "Notify status updated successfully",
	}
}

func GetProductSizesAndStock(db *gorm.DB, productId int) WishlistModel.ProductSizeStockResponse {
	log := logger.InitLogger()

	var result []WishlistModel.ProductSizeStock

	fmt.Println("productId", productId)
	err := db.Raw(WishlistQuery.GetProductSizesAndStockQuery, productId).Scan(&result).Error
	if err != nil {
		log.Errorf("Error fetching sizes and stock: %v", err)
		return WishlistModel.ProductSizeStockResponse{
			Status:  false,
			Message: "Failed to fetch product sizes and stock",
			Data:    nil,
		}
	}

	fmt.Println("result", result)

	return WishlistModel.ProductSizeStockResponse{
		Status:  true,
		Message: "Product sizes and stock fetched successfully",
		Data:    result,
	}
}