package getProductsModel

// Response for the API
type ListBrandsWithLogoOptionsResp struct {
	Status       bool                    `json:"status"`
	Message      string                  `json:"message"`
	BrandOptions []BrandOptionsQueryResp `json:"parentCategoryOptions,omitempty"`
}

type BrandOptionsQueryResp struct {
	RefApplicationId     int    `json:"refApplicationId" gorm:"column:refApplicationId"`
	RefBrandName         string `json:"refBrandName" gorm:"column:refBrandName"`
	RefLogoFileName      string `json:"refLogo" gorm:"column:refLogo"`
	RefCelebrityFileName string `json:"refCelebrityImagePath" gorm:"column:refCelebrityImagePath"`
}

type ParentCategoryResp struct {
	RefParentCateroryMapId     int    `json:"refParentCateroryMapId" gorm:"column:refParentCateroryMapId"`
	RefCategoryGenderId        int    `json:"refCategoryGenderId" gorm:"column:refCategoryGenderId"`
	GenderName                 string `json:"genderName" gorm:"column:genderName"`
	ParentCategoryName         string `json:"parentCategoryName" gorm:"column:parentCategoryName"`
	RefParentCategoryImagePath string `json:"refParentCategoryImagePath" gorm:"column:refParentCategoryImagePath"`
	RefIfVisible               bool   `json:"refIfVisible" gorm:"column:refIfVisible"`
}

type ListParentCategoryResp struct {
	Status           bool                 `json:"status"`
	Message          string               `json:"message"`
	ParentCategories []ParentCategoryResp `json:"parentCategories"`
}

// For scanning parent category rows
type ParentCategoryRaw struct {
	RefParentCateroryMapId     int    `gorm:"column:refParentCateroryMapId"`
	RefCategoryGenderId        int    `gorm:"column:refCategoryGenderId"`
	GenderName                 string `gorm:"column:genderName"`
	ParentCategoryId           int    `gorm:"column:parentCategoryId"`
	ParentCategoryName         string `gorm:"column:parentCategoryName"`
	RefParentCategoryImagePath string `gorm:"column:refParentCategoryImagePath"`
	ParentVisible              bool   `gorm:"column:parentVisible"`
}

// Final struct with subcategories populated manually
type SubCategory struct {
	SubCategoryId   int    `gorm:"column:refSubCategoryId"`
	SubCategoryName string `gorm:"column:subCategoryName"`
}
type ParentCategoryWithSub struct {
	RefParentCateroryMapId     int           `json:"refParentCateroryMapId"`
	RefCategoryGenderId        int           `json:"refCategoryGenderId"`
	GenderName                 string        `json:"genderName"`
	ParentCategoryId           int           `json:"parentCategoryId"`
	ParentCategoryName         string        `json:"parentCategoryName"`
	RefParentCategoryImagePath string        `json:"refParentCategoryImagePath"`
	ParentVisible              bool          `json:"parentVisible"`
	SubCategories              []SubCategory `json:"subCategories"`
}

type ListParentCategoryWithSubResp struct {
	Status           bool                    `json:"status"`
	Message          string                  `json:"message"`
	ParentCategories []ParentCategoryWithSub `json:"parentCategories"`
}

type ProductDefaultCatalogResponseOptionsResp struct {
	Status         bool                            `json:"status"`
	Message        string                          `json:"message"`
	ProductOptions []ProductDefaultCatalogResponse `json:"ProductOptions,omitempty"`
}

type ProductDefaultCatalogResponse struct {
	RefProductId               int     `json:"refProductId" gorm:"column:refProductId"`
	RefProductName             string  `json:"refProductName" gorm:"column:refProductName"`
	RefProductMrp              float64 `json:"refProductMrp" gorm:"column:refProductMrp"`
	RefProductMsp              float64 `json:"refProductMsp" gorm:"column:refProductMsp"`
	RefProductDetailAngleImage string  `json:"refProductDetailAngleImage" gorm:"column:refProductDetailAngleImage"`
	RefBrandName               string  `json:"refBrandName" gorm:"column:refBrandName"`
	RefParentCategoryName      string  `json:"refParentCategoryName" gorm:"column:refParentCategoryName"`
	RefSubCategory             string  `json:"refSubCategory" gorm:"column:refSubCategory"`
	OfferPercentage            string  `json:"offerPercentage" gorm:"column:offerPercentage"`
}

type ListProductsOptions struct {
	Status   bool           `json:"status"`
	Message  string         `json:"message"`
	Products []ProductsResp `json:"ProductsResp"`
}

type ProductsResp struct {
	RefProductId                        int    `json:"refProductId" gorm:"column:refProductId"`
	RefProductName                      string `json:"refProductName" gorm:"column:refProductName"`
	RefProductMrp                       string `json:"refProductMrp" gorm:"column:refProductMrp"`
	RefProductDetailAngleImage          string `json:"refProductDetailAngleImage" gorm:"column:refProductDetailAngleImage"`
	RefProductFrontImage                string `json:"refProductFrontImage" gorm:"column:refProductFrontImage"`
	RefProductSideImage                 string `json:"refProductSideImage" gorm:"column:refProductSideImage"`
	RefProductBackImage                 string `json:"refProductBackImage" gorm:"column:refProductBackImage"`
	RefProductAdditionalImage           string `json:"refProductAdditionalImage" gorm:"column:refProductAdditionalImage"`
	RefProductVideo                     string `json:"refProductVideo" gorm:"column:refProductVideo"`
	RefProductMsp                       string `json:"refProductMsp" gorm:"column:refProductMsp"`
	RefBrandName                        string `json:"refBrandName" gorm:"column:refBrandName"`
	OfferPercentage                     string `json:"offerPercentage" gorm:"column:offerPercentage"`
	SoldBy                              string `json:"soldBy" gorm:"column:soldBy"`
	RefProductCountryOfOrigin           string `json:"refProductCountryOfOrigin" gorm:"column:refProductCountryOfOrigin"`
	RefProductManufactureNameAndAddress string `json:"refProductManufactureNameAndAddress" gorm:"column:refProductManufactureNameAndAddress"`
}
