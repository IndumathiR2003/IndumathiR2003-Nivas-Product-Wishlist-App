package CartQuery

var AddToCartQuery = `
INSERT INTO public."refUserCart" (
  "refProductId",
  "refQuantity",
  "refWishlistId",
  "refUserId",
  "refCartStatus",
  "refCreateAT",
  "refCreateBy",
  "refIsDelete"
)
VALUES ($1, $2, $3, $4, 'ACTIVE', $5, $6, FALSE)
RETURNING "refCartId";
`

var GetUserCartQuery = `
SELECT 
  c."refCartId" AS "RefCartId",
  c."refQuantity" AS "RefQuantity",
  c."refWishlistId" AS "RefWishlistId",
  p."refProductId" AS "RefProductId",
  p."refProductName" AS "RefProductName",
  rpdc."refProductAttributeValue" AS "RefProductSize",
  p."refProductMrp" AS "RefProductMrp",
  p."refProductMsp" AS "RefProductMsp",
  p."refProductDetailAngleImage" AS "RefProductDetailAngleImage",
  b."refBrandName" AS "RefBrandName",
  s."refSubCategory" AS "RefSubCategory",
  (
    SELECT COALESCE(SUM(rpc."refTotalStock"), 0)
    FROM product."refProductCatLog" rpc
    WHERE rpc."refProductColorGroupId" = (
      SELECT "refProductColorGroupId"
      FROM product."refProductCatLog"
      WHERE "refProductId" = p."refProductId"
    )
  ) AS "RefTotalStock"
FROM public."refUserCart" c
JOIN product."refproductDefaultCatLog" p 
  ON c."refProductId" = p."refProductId"
JOIN brand."refBrandApplication" b 
  ON b."refApplicationId" = p."refBrandId"
JOIN productcategory."refSubCategory" s 
  ON s."refSubCategoryId" = p."refSubCategoryId"
JOIN product."refProductDynamicCatLog" rpdc ON rpdc."refProductId" = p."refProductId"
WHERE c."refCartStatus" = 'ACTIVE'
  AND c."refUserId" = $1
AND c."refIsDelete" = FALSE
ORDER BY c."refCartId"
`

var UpdateCartQuery = `
UPDATE public."refUserCart"
SET 
  "refProductId" = $1,
  "refQuantity" = $2,
  "refUpdateAt" = $3,
  "refUpdateBy" = $4
WHERE "refCartId" = $5
  AND "refIsDelete" = FALSE
RETURNING *;
`
