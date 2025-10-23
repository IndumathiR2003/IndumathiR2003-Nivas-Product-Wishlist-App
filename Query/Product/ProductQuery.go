package getProductsQuery

var GetBrandListQuery = `
SELECT
  bd."refBrandName",
  bd."refApplicationId",
  rd."refLogo",
  rd."refCelebrityImagePath"
FROM
  brand."refBrandApplication" bd
LEFT JOIN brand."refDocuments" rd
  ON CAST(bd."refDocumentsId" AS INTEGER) = CAST(rd."refDocumentsId" AS INTEGER)
WHERE
  bd."refApplicationStatus" = 4;
`
var GetParentCategorybyGenderQuery = `
SELECT
  pcm."refParentCateroryMapId",
  g."refCategoryGenderId",
  g."refCategoryGenderName" AS "genderName",
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
WHERE
  pcm."refIfVisible" = TRUE AND g."refCategoryGenderId" = $1
ORDER BY
  g."refCategoryGenderName",
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
  br."refBrandName",
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
  pr."refProductManufactureNameAndAddress"
FROM
  product."refproductDefaultCatLog" pr
  LEFT JOIN brand."refBrandApplication" br ON CAST(br."refApplicationId" AS INTEGER) = CAST(pr."refBrandId" AS INTEGER)
  LEFT JOIN productcategory."refParentCategory" pc ON CAST(pc."refParentCategoryId" AS INTEGER) = CAST(pr."refParentCategoryId" AS INTEGER)
  LEFT JOIN productcategory."refSubCategory" sc ON CAST(sc."refSubCategoryId" AS INTEGER) = CAST(pr."refSubCategoryId" AS INTEGER)
WHERE
  pr."refProductId" = $1;
`