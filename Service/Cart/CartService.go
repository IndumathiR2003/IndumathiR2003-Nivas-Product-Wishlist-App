package CartService

import (
	"fmt"
	logger "nivasProductBackendApp/Helper/Logger"
	CartModel "nivasProductBackendApp/Model/Cart"
	CartQuery "nivasProductBackendApp/Query/Cart"
	minioService "nivasProductBackendApp/Helper/MinIo"
	"gorm.io/gorm"
)

func AddToCart(db *gorm.DB, payload CartModel.AddToCartRequest) CartModel.AddToCartResponse {
	log := logger.InitLogger()

	var cartId int
	err := db.Raw(CartQuery.AddToCartQuery,
		payload.RefProductId,
		payload.RefQuantity,
		payload.RefWishlistId,
		payload.RefUserId,
		"2025",
		payload.RefUserId,
	).Scan(&cartId).Error

	if err != nil {
		log.Errorf("Error adding to cart: %v", err)
		return CartModel.AddToCartResponse{
			Status:  false,
			Message: "Failed to add to cart",
		}
	}

	fmt.Println("Cart ID:", cartId)

	return CartModel.AddToCartResponse{
		Status:  true,
		Message: "Product added to cart successfully",
		CartId:  cartId,
	}
}

func ListUserCart(db *gorm.DB, userId int) CartModel.CartListResponse {
	log := logger.InitLogger()
	var cartItems []CartModel.CartItem

	err := db.Raw(CartQuery.GetUserCartQuery, userId).Scan(&cartItems).Error
	if err != nil {
		log.Errorf("Error fetching user cart: %v", err)
		return CartModel.CartListResponse{
			Status:  false,
			Message: "Failed to fetch cart products",
			Data:    nil,
		}
	}

	fmt.Println("Cart items:", cartItems)

	// 🔹 Add hardcoded fields
	for i := range cartItems {
		if cartItems[i].RefProductDetailAngleImage != "" {
            url, err := minioService.GetFileURL(cartItems[i].RefProductDetailAngleImage, 60)
            if err == nil {
                cartItems[i].RefProductDetailAngleImage = url
            }
        }
	}

	return CartModel.CartListResponse{
		Status:  true,
		Message: "Cart products fetched successfully",
		Data:    cartItems,
	}
}

func UpdateCart(db *gorm.DB, req CartModel.UpdateCartRequest) CartModel.UpdateCartResponse {
	log := logger.InitLogger()
	var updatedItem CartModel.CartItem

	err := db.Raw(CartQuery.UpdateCartQuery, req.RefProductId, req.RefQuantity, "2025", req.RefUpdateBy, req.RefCartId).
		Scan(&updatedItem).Error

	if err != nil {
		log.Errorf("Error updating cart: %v", err)
		return CartModel.UpdateCartResponse{
			Status:  false,
			Message: "Failed to update cart item",
		}
	}

	return CartModel.UpdateCartResponse{
		Status:  true,
		Message: "Cart item updated successfully",
		Data:    updatedItem,
	}
}
