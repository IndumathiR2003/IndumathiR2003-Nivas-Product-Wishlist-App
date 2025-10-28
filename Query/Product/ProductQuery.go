package getProductsQuery

var GetBrandListQuery = `
SELECT
  pm."refPaidBrandMappingId",
  br."refBrandName",
  d."refLogo",
  d."refCelebrityImagePath",
  pm."isFeatured",
  pm."isHomePageAllowed"
FROM
  brand."refPaidBrandMapping" pm
  LEFT JOIN brand."refBrandApplication" br ON br."refApplicationId"::INTEGER = pm."refBrandId"::INTEGER
  LEFT JOIN brand."refDocuments" d ON d."refDocumentsId"::INTEGER = br."refDocumentsId"::INTEGER
`
var GetParentCategorybyGenderQuery = `
SELECT DISTINCT ON (pc."refParentCategoryId")
  g."refCategoryGenderId",
  g."refCategoryGenderName" AS "genderName",
  pcm."refParentCateroryMapId",
  pc."refParentCategoryName" AS "parentCategoryName",
  pcm."refParentCategoryImagePath",
  pcm."refIfVisible"
FROM
  productcategory."refParentCategoryMapping" pcm
JOIN
  productcategory."refCategoryGender" g
  ON pcm."refCategoryGenderId" = g."refCategoryGenderId"
JOIN
  productcategory."refParentCategory" pc
  ON pcm."refParentCategoryId" = pc."refParentCategoryId"
JOIN
  product."refproductDefaultCatLog" pdc
  ON pdc."refParentCategoryId" = pc."refParentCategoryId"
WHERE
  pcm."refIfVisible" = TRUE AND g."refCategoryGenderId" = $1
ORDER BY
  pc."refParentCategoryId",
  pc."refParentCategoryName";
`
var ListParentandSubCategoriesQuery = `
SELECT
    pcm."refParentCateroryMapId",
    g."refCategoryGenderId",
    g."refCategoryGenderName" AS "genderName",
    pc."refParentCategoryId",
    pc."refParentCategoryName" AS "parentCategoryName",
    pc."refParentCategoryImagePath",
    pcm."refIfVisible" AS parentVisible,
    json_agg(
        json_build_object(
            'subCategoryId', sc."refSubCategoryId",
            'subCategoryName', sc."refSubCategory"
        )
    ) AS subCategories
FROM productcategory."refParentCategoryMapping" pcm
JOIN productcategory."refCategoryGender" g
    ON pcm."refCategoryGenderId" = g."refCategoryGenderId"
JOIN productcategory."refParentCategory" pc
    ON pcm."refParentCategoryId" = pc."refParentCategoryId"
LEFT JOIN productcategory."refParentSubCategoryMapping" pscm
    ON pcm."refParentCateroryMapId" = pscm."refParentGenderMappingId"
LEFT JOIN productcategory."refSubCategory" sc
    ON pscm."refSubCategoryId" = sc."refSubCategoryId" 
    AND sc."refIfVisible" = TRUE
WHERE pcm."refIfVisible" = TRUE
GROUP BY
    pcm."refParentCateroryMapId",
    g."refCategoryGenderId",
    g."refCategoryGenderName",
    pc."refParentCategoryId",
    pc."refParentCategoryName",
    pc."refParentCategoryImagePath",
    pcm."refIfVisible"
ORDER BY
    g."refCategoryGenderName",
    pc."refParentCategoryName";
`

var GetVisibleProductCatalogQuery = `
SELECT
  pr."refProductId",
  pr."refProductName",
  pr."refProductMrp",
  pr."refProductDetailAngleImage",
  pr."refProductMsp",
  br."refBrandName",
  pc."refParentCategoryName",
  sc."refSubCategory",
  CONCAT(
    ROUND(
      ((CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)) 
        / CAST(pr."refProductMrp" AS NUMERIC)) * 100
    )::INT,
    '%'
  ) AS "offerPercentage"
FROM
  product."refproductDefaultCatLog" pr
  LEFT JOIN brand."refBrandApplication" br 
    ON CAST(br."refApplicationId" AS INTEGER) = CAST(pr."refBrandId" AS INTEGER)
  LEFT JOIN productcategory."refParentCategory" pc 
    ON CAST(pc."refParentCategoryId" AS INTEGER) = CAST(pr."refParentCategoryId" AS INTEGER)
  LEFT JOIN productcategory."refSubCategory" sc 
    ON CAST(sc."refSubCategoryId" AS INTEGER) = CAST(pr."refSubCategoryId" AS INTEGER)
WHERE 
  pc."refIfVisible" IS TRUE 
  AND sc."refIfVisible" IS TRUE;
`

var GetproductDetailsQuery = `
SELECT
  pr."refProductId",
  br."refBrandName",
  pr."refProductName",
  pr."refProductDescription",
  pr."refProductMrp",
  pr."refProductDetailAngleImage",
  pr."refProductFrontImage",
  pr."refProductSideImage",
  pr."refProductBackImage",
  pr."refProductLookShotImage",
  pr."refProductAdditionalImage",
  pr."refProductVideo",
  pr."refProductMsp",
 
  pr."refProductWarrantyAndReturnPolicy",
  CONCAT(
    ROUND(
      (
        (
          CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)
        ) / CAST(pr."refProductMrp" AS NUMERIC)
      ) * 100
    )::INT,
    '%'
  ) AS "offerPercentage",
  br."refBrandName" AS "soldBy",
  pr."refProductCountryOfOrigin",
  pr."refProductManufactureNameAndAddress",
  ROUND(AVG(CAST(ur."refRatingCount" AS NUMERIC)), 1) AS "averageRating",

  COUNT(ur."refUserReviewId") AS "totalReviews"
FROM
  product."refproductDefaultCatLog" pr
  LEFT JOIN brand."refBrandApplication" br ON CAST(br."refApplicationId" AS INTEGER) = CAST(pr."refBrandId" AS INTEGER)
  LEFT JOIN productcategory."refParentCategory" pc ON CAST(pc."refParentCategoryId" AS INTEGER) = CAST(pr."refParentCategoryId" AS INTEGER)
  LEFT JOIN productcategory."refSubCategory" sc ON CAST(sc."refSubCategoryId" AS INTEGER) = CAST(pr."refSubCategoryId" AS INTEGER)
  LEFT JOIN public."refUserReviews" ur ON CAST(ur."refProductId" AS INTEGER) = CAST(pr."refProductId" AS INTEGER)
WHERE
  pr."refProductId" = $1
GROUP BY pr."refProductId",
br."refApplicationId";
`

var GetProductsByGender = `
SELECT
  pr."refProductId",
  pr."refProductName",
  pr."refProductDescription",
  pr."refProductMrp",
  pr."refProductDetailAngleImage",
  pr."refProductMsp",
  br."refBrandName",
  g."refCategoryGenderName",
  pc."refParentCategoryName",
  CONCAT(
    ROUND(
      (
        (
          CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)
        ) / CAST(pr."refProductMrp" AS NUMERIC)
      ) * 100
    )::INT,
    '%'
  ) AS "offerPercentage"
FROM product."refproductDefaultCatLog" pr
LEFT JOIN brand."refBrandApplication" br ON br."refApplicationId"::INTEGER = pr."refBrandId"::INTEGER
LEFT JOIN productcategory."refParentCategory" pc ON pc."refParentCategoryId"::INTEGER = pr."refParentCategoryId"::INTEGER
LEFT JOIN productcategory."refCategoryGender" g ON g."refCategoryGenderId"::INTEGER = pr."refGenderId"::INTEGER
WHERE pr."refGenderId" = ?;
`
var GetProductsByParentCategory = `
SELECT
  pr."refProductId",
  pr."refProductName",
  pr."refProductDescription",
  pr."refProductMrp",
  pr."refProductDetailAngleImage",
  pr."refProductMsp",
  br."refBrandName",
  g."refCategoryGenderName",
  pc."refParentCategoryName",
  CONCAT(
    ROUND(
      (
        (
          CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)
        ) / CAST(pr."refProductMrp" AS NUMERIC)
      ) * 100
    )::INT,
    '%'
  ) AS "offerPercentage"
FROM product."refproductDefaultCatLog" pr
LEFT JOIN brand."refBrandApplication" br ON br."refApplicationId"::INTEGER = pr."refBrandId"::INTEGER
LEFT JOIN productcategory."refParentCategory" pc ON pc."refParentCategoryId"::INTEGER = pr."refParentCategoryId"::INTEGER
LEFT JOIN productcategory."refCategoryGender" g ON g."refCategoryGenderId"::INTEGER = pr."refGenderId"::INTEGER
WHERE pr."refParentCategoryId" = ?;
`

var ListCategoryPageParentCategroryQuery = `
SELECT DISTINCT ON (pc."refParentCategoryId")
  cg."refCategoryGenderId",
  cg."refCategoryGenderName",
  cg."refGenderImagePath",
  pc."refParentCategoryId",
  pc."refParentCategoryName",
  pc."refCategoryImagePath"
FROM
  productcategory."refCategoryGender" cg
  LEFT JOIN productcategory."refParentCategoryMapping" pcm 
    ON pcm."refCategoryGenderId"::INTEGER = cg."refCategoryGenderId"::INTEGER
  LEFT JOIN productcategory."refParentCategory" pc 
    ON pc."refParentCategoryId"::INTEGER = pcm."refParentCategoryId"::INTEGER
WHERE
  pc."refParentCategoryId" IS NOT NULL
ORDER BY
  pc."refParentCategoryId",
  cg."refCategoryGenderName";
`

var ListParentCategrorywithSubcategoryQuery = `
SELECT
  cg."refCategoryGenderId",
  cg."refCategoryGenderName",
  pc."refParentCategoryId",
  pc."refParentCategoryName",
  sc."refSubCategoryId",
  sc."refSubCategory"
FROM
  productcategory."refCategoryGender" cg
  LEFT JOIN productcategory."refParentCategoryMapping" pcm ON pcm."refCategoryGenderId"::INTEGER = cg."refCategoryGenderId"::INTEGER
  LEFT JOIN productcategory."refParentCategory" pc ON pc."refParentCategoryId"::INTEGER = pcm."refParentCategoryId"::INTEGER
  LEFT JOIN productcategory."refParentSubCategoryMapping" scm ON scm."refParentGenderMappingId"::INTEGER = pcm."refParentCateroryMapId"::INTEGER
  LEFT JOIN productcategory."refSubCategory" sc ON sc."refSubCategoryId"::INTEGER = scm."refSubCategoryId"::INTEGER
  WHERE cg."refCategoryGenderId"=$1;
  `
var ListSubcategoryQuery = `
SELECT
  cg."refCategoryGenderId",
  cg."refCategoryGenderName",
  pc."refParentCategoryId",
  pc."refParentCategoryName",
  sc."refSubCategoryId",
  sc."refSubCategory"
FROM
  productcategory."refCategoryGender" cg
  LEFT JOIN productcategory."refParentCategoryMapping" pcm ON pcm."refCategoryGenderId"::INTEGER = cg."refCategoryGenderId"::INTEGER
  LEFT JOIN productcategory."refParentCategory" pc ON pc."refParentCategoryId"::INTEGER = pcm."refParentCategoryId"::INTEGER
  LEFT JOIN productcategory."refParentSubCategoryMapping" scm ON scm."refParentGenderMappingId"::INTEGER = pcm."refParentCateroryMapId"::INTEGER
  LEFT JOIN productcategory."refSubCategory" sc ON sc."refSubCategoryId"::INTEGER = scm."refSubCategoryId"::INTEGER
  WHERE pc."refParentCategoryId"=$1;
  
  `
var ListProductsByParentCategory = `
SELECT
  pr."refProductId",
  br."refBrandName",
  pr."refProductName",
  pr."refProductMrp",
  pc."refParentCategoryName",
  pr."refProductDetailAngleImage",
  pr."refProductMsp",
  sc."refSubCategoryId",
  sc."refSubCategory",
  CASE
    WHEN pr."refProductMrp" > pr."refProductMsp" THEN CONCAT(
      ROUND(
        (
          (
            CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)
          ) / CAST(pr."refProductMrp" AS NUMERIC)
        ) * 100
      )::INT,
      '%'
    )
    ELSE NULL
  END AS "offerPercentage"
FROM
  productcategory."refParentCategory" pc
  LEFT JOIN product."refproductDefaultCatLog" pr ON pr."refParentCategoryId"::INTEGER = pc."refParentCategoryId"::INTEGER
  LEFT JOIN brand."refBrandApplication" br ON br."refApplicationId"::INTEGER = pr."refBrandId"::INTEGER
  LEFT JOIN productcategory."refSubCategory" sc ON sc."refSubCategoryId"::INTEGER = pr."refSubCategoryId"::INTEGER
WHERE
  pc."refParentCategoryId" = $1;
`

var ListProductsBySubCategory = `
SELECT
  pr."refProductId",
  br."refBrandName",
  pr."refProductName",
  pr."refProductMrp",
  sc."refSubCategory",
  pr."refProductDetailAngleImage",
  pr."refProductMsp",
  CASE
    WHEN pr."refProductMrp" > pr."refProductMsp" THEN CONCAT(
      ROUND(
        (
          (
            CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)
          ) / CAST(pr."refProductMrp" AS NUMERIC)
        ) * 100
      )::INT,
      '%'
    )
    ELSE NULL
  END AS "offerPercentage"
FROM
  productcategory."refSubCategory" sc
  JOIN product."refproductDefaultCatLog" pr ON sc."refSubCategoryId"::INTEGER = pr."refSubCategoryId"::INTEGER
  JOIN brand."refBrandApplication" br ON br."refApplicationId"::INTEGER = pr."refBrandId"::INTEGER
WHERE
  sc."refSubCategoryId" = $1;
`

var ListNewArrivalsQuery = `
SELECT
  pr."refProductId",
  pr."refProductName",
  pr."refProductMrp",
  pr."refProductDetailAngleImage",
  pr."refProductMsp",
  pr."refCreateAt",
  br."refBrandName",
  br."refApplicationId",
  d."refLogo",
  pc."refParentCategoryName",
  sc."refSubCategory",
  CONCAT(
    ROUND(
      (
        (
          CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)
        ) / CAST(pr."refProductMrp" AS NUMERIC)
      ) * 100
    )::INT,
    '%'
  ) AS "offerPercentage"
FROM
  product."refproductDefaultCatLog" pr
  LEFT JOIN brand."refBrandApplication" br 
    ON CAST(br."refApplicationId" AS INTEGER) = CAST(pr."refBrandId" AS INTEGER)
  LEFT JOIN brand."refDocuments" d 
    ON CAST(d."refDocumentsId" AS INTEGER) = CAST(br."refDocumentsId" AS INTEGER)
  LEFT JOIN productcategory."refParentCategory" pc 
    ON CAST(pc."refParentCategoryId" AS INTEGER) = CAST(pr."refParentCategoryId" AS INTEGER)
  LEFT JOIN productcategory."refSubCategory" sc 
    ON CAST(sc."refSubCategoryId" AS INTEGER) = CAST(pr."refSubCategoryId" AS INTEGER)
  LEFT JOIN brand."refPaidBrandMapping" pbm 
    ON CAST(pbm."refBrandId" AS INTEGER) = CAST(pr."refBrandId" AS INTEGER)
WHERE
  pbm."isHomePageAllowed" IS TRUE
ORDER BY
  pr."refCreateAt" DESC;
`
var GetNewArrivalProductsByBrandQuery = `
SELECT
  pr."refProductId",
  pr."refProductName",
  pr."refProductMrp",
  pr."refProductDetailAngleImage",
  pr."refProductMsp",
  pr."refCreateAt",
  br."refBrandName",
  br."refApplicationId",
  d."refLogo",
  pc."refParentCategoryName",
  sc."refSubCategory",
    sc."refSubCategoryId",

  CONCAT(
    ROUND(
      (
        (
          CAST(pr."refProductMrp" AS NUMERIC) - CAST(pr."refProductMsp" AS NUMERIC)
        ) / CAST(pr."refProductMrp" AS NUMERIC)
      ) * 100
    )::INT,
    '%'
  ) AS "offerPercentage"
FROM
  product."refproductDefaultCatLog" pr
  LEFT JOIN brand."refBrandApplication" br 
    ON CAST(br."refApplicationId" AS INTEGER) = CAST(pr."refBrandId" AS INTEGER)
  LEFT JOIN brand."refDocuments" d 
    ON CAST(d."refDocumentsId" AS INTEGER) = CAST(br."refDocumentsId" AS INTEGER)
  LEFT JOIN productcategory."refParentCategory" pc 
    ON CAST(pc."refParentCategoryId" AS INTEGER) = CAST(pr."refParentCategoryId" AS INTEGER)
  LEFT JOIN productcategory."refSubCategory" sc 
    ON CAST(sc."refSubCategoryId" AS INTEGER) = CAST(pr."refSubCategoryId" AS INTEGER)
  LEFT JOIN brand."refPaidBrandMapping" pbm 
    ON CAST(pbm."refBrandId" AS INTEGER) = CAST(pr."refBrandId" AS INTEGER)
WHERE
  pbm."refBrandId" = $1
ORDER BY
  pr."refCreateAt" DESC;
`