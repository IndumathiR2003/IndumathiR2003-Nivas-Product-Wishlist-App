package WishlistQuery

var GetWishlistProductsQuery = `
SELECT
  w."refWishlistId",
  p."refProductId",
  b."refBrandName",
  p."refProductName",
  p."refProductMrp",
  p."refProductMsp",
  p."refProductDetailAngleImage",
  s."refSubCategoryId",
  s."refSubCategory",
  w."refNotify",
  COALESCE((
    SELECT SUM(rpc."refTotalStock")
    FROM product."refProductCatLog" rpc
    WHERE rpc."refProductColorGroupId" = (
      SELECT "refProductColorGroupId"
      FROM product."refProductCatLog"
      WHERE "refProductId" = p."refProductId"
    )
  ), 0) AS "refTotalStock"
FROM public."refProductWishlist" AS w
JOIN public."Users" AS u
  ON u."refUserId" = w."refUserId"
JOIN product."refproductDefaultCatLog" AS p
  ON p."refProductId" = w."refProductId"
JOIN brand."refBrandApplication" AS b
  ON b."refApplicationId" = p."refBrandId"
JOIN productcategory."refSubCategory" AS s
  ON s."refSubCategoryId" = p."refSubCategoryId"
LEFT JOIN product."refProductCatLog" rpc
  ON rpc."refProductId" = p."refProductId"
WHERE u."refUserId" = $1
  AND w."refWishlistStatus" = TRUE
ORDER BY w."refWishlistId";
`

var AddToWishlistQuery = `
INSERT INTO public."refProductWishlist" 
("refUserId", "refProductId", "refWishlistStatus", "refCreateBy", "refCreateAt")
VALUES ($1, $2, TRUE, $3, $4)
ON CONFLICT ("refUserId", "refProductId")
DO UPDATE SET "refWishlistStatus" = TRUE
RETURNING "refWishlistId";
`

var RemoveFromWishlistQuery = `
UPDATE public."refProductWishlist"
SET "refWishlistStatus" = FALSE
WHERE "refWishlistId" = ?
RETURNING "refWishlistId";
`

var UpdateNotifyStatusQuery = `
UPDATE public."refProductWishlist"
SET "refNotify" = TRUE
WHERE "refWishlistId" = $1
`

var GetProductSizesAndStockQuery = `
SELECT 
  rpdc."refProductAttributeValue",
  rpc."refTotalStock",
  rpc."refProductId"
FROM product."refProductDynamicCatLog" rpdc
JOIN productcategory."refProductDynamicAttributes" rpda 
  ON rpda."refDynamicAttributeId" = rpdc."refProductAttributeId"
JOIN product."refProductCatLog" rpc 
  ON rpc."refProductId" = rpdc."refProductId"
WHERE rpc."refProductColorGroupId" = (
  SELECT p."refProductColorGroupId"
  FROM product."refProductCatLog" p
  WHERE p."refProductId" = $1
)`