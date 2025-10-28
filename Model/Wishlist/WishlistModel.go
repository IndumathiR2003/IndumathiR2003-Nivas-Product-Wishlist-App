package WishlistModel

type ListWishlistModel struct {
    RefWishlistId int     `json:"refWishlistId" gorm:"column:refWishlistId"`
    RefProductId  int     `json:"refProductId" gorm:"column:refProductId"`
    RefBrandName  string  `json:"refBrandName" gorm:"column:refBrandName"`
    RefProductName string `json:"refProductName" gorm:"column:refProductName"`
	RefProductImage string `json:"refProductImage" gorm:"column:refProductDetailAngleImage"`
    RefProductMrp  float64 `json:"refProductMrp" gorm:"column:refProductMrp"`
    RefProductMsp  float64 `json:"refProductMsp" gorm:"column:refProductMsp"`
	// 💡 Hardcoded / Derived fields
    RefRatingStar          float64 `json:"refRatingStar"`
    RefRatingTotalReviewers int     `json:"refRatingTotalReviewers"`
    RefSubCategoryId       int     `json:"refSubCategoryId" gorm:"column:refSubCategoryId"`
    RefSubCategory         string  `json:"refSubCategoryName" gorm:"column:refSubCategory"`
    RefTotalStock          int     `json:"refTotalStock" gorm:"column:refTotalStock"`
    RefProductNotify       bool    `json:"refProductNotify" gorm:"column:refNotify"`
}


type ListWishlistModelResponse struct {
	Status   bool                `json:"status"`
	Message  string              `json:"message"`
	Products []ListWishlistModel `json:"wishlist,omitempty"`
}

type AddWishlistRequest struct {
	UserId    int `json:"userId" binding:"required"`
	ProductId int `json:"productId" binding:"required"`
}

type AddWishlistResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type RemoveWishlistRequest struct {
	WishlistId int `json:"wishlistId" binding:"required"`
}

type RemoveWishlistResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type WishlistUpdateResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ProductSizeStock struct {
	RefProductAttributeValue string `json:"refProductAttributeValue" gorm:"column:refProductAttributeValue"`
	RefTotalStock            int    `json:"refTotalStock" gorm:"column:refTotalStock"`
	RefProductId             int    `json:"refProductId" gorm:"column:refProductId"`
}

type ProductSizeStockResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Data    []ProductSizeStock `json:"data"`
}