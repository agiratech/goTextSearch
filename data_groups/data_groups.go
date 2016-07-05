package data_groups

import (
  _"fmt"
  "github.com/agiratech/goTextSearch/config"
  "github.com/agiratech/goTextSearch/common_struct"
)

func BrandClassification(addrProductInfo *common_struct.ProductInfo){
  query := `SELECT
    COALESCE(PROD.retailer_product_configuration_id, 0) AS retailer_prod_conf_id,
    COALESCE(to_char(OCC.imported_at, 'YYYYMMDD'), '') AS date,
    COALESCE(to_char(OCC.imported_at, 'Dy'), '') AS day_of_the_week,
    COALESCE(RET.id, 0) AS retailer_id,
    COALESCE(RET.name, '') AS retailer_name,
    COALESCE(PROD.id, 0) AS product_id,
    COALESCE(PRO_DET.id, 0) AS product_detail_id,
    COALESCE(PRO_DET.product_url, '') AS product_url,
    COALESCE(PROD.brand, '') AS brand,
    COALESCE(PROD.category, '') AS category,
    COALESCE(PRO_DET.model_number, '') AS model_number,
    COALESCE(PRO_DET.product_code, '') AS product_code,
    COALESCE(PRO_DET.sku_code, '') AS sku_code,
    LOWER(COALESCE(PROD.title, '')) AS product_name,
    regexp_replace(COALESCE(PROD.description, ''), E'[\n\r]+', ' ', 'g' ) AS product_description,
    COALESCE(PRO_DET.barcode, '') AS barcode,
    COALESCE(PRO_DET.color, '') AS color,
    COALESCE(PRO_DET.status, '') AS status,
    COALESCE(PRO_DET.size, '') AS size,
    COALESCE(PROD.gender, '') AS gender,
    COALESCE(PRO_DET.no_of_gears, '') AS no_of_gears,
    COALESCE(PRO_DET.suspension, '') AS suspension,
    COALESCE(PRO_DET.thumb_image_url, '') AS image_url_thumbnail,
    COALESCE(PRO_DET.preview_image_url  , '') AS image_url_preview,
    COALESCE(RET.currency, '') AS currency,
    COALESCE(PRO_DET.original_price, 0) AS original_price,
    COALESCE(PRO_DET.selling_price, 0) AS sale_price
    FROM occurences AS OCC
    JOIN product_details AS PRO_DET ON PRO_DET.id = OCC.product_detail_id
    JOIN products AS PROD ON PROD.id = PRO_DET.product_id
    JOIN retailers AS RET ON RET.id = PROD.retailer_id
    where PROD.brand LIKE $1`
  _, _ = config.Dbmap.Select(&addrProductInfo.Products, query, addrProductInfo.TargetBrand)
   // return addrProductInfo
}