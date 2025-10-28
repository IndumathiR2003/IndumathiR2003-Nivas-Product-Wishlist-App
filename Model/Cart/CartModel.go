package CartModel

type AddToCartRequest struct {
	RefProductId  int  `json:"refProductId"`
	RefQuantity   int  `json:"refQuantity"`
	RefWishlistId *int `json:"refWishlistId,omitempty"` // can be null
	RefUserId     int  `json:"refUserId"`
}

type AddToCartResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	CartId  int    `json:"cartId,omitempty"`
}

type CartItem struct {
	RefCartId                int     `json:"refCartId"`
	RefQuantity              int     `json:"refQuantity"`
	RefWishlistId            *int    `json:"refWishlistId,omitempty"`
	RefProductId             int     `json:"refProductId"`
	RefProductName           string  `json:"refProductName"`
	RefProductSize string  `json:"refProductAttributeValue"`
	RefProductMrp            float64 `json:"refProductMrp"`
	RefProductMsp            float64 `json:"refProductMsp"`
	RefProductDetailAngleImage string `json:"refProductDetailAngleImage"`
	RefBrandName             string  `json:"refBrandName"`
	RefSubCategory           string  `json:"refSubCategory"`
	RefTotalStock            int     `json:"refTotalStock"`
}


type CartListResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []CartItem `json:"data"`
}

type UpdateCartRequest struct {
	RefProductId  int  `json:"refProductId"`
	RefCartId   int `json:"refCartId"`
	RefQuantity int `json:"refQuantity"`
	RefUpdateBy int `json:"refUserId"`
}

type UpdateCartResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}