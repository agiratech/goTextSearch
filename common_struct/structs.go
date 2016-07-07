package common_struct

type ProductsData struct{
  RetailerProductConfigurationId string `db:"retailer_prod_conf_id"`
  Date string `db:"date"`
  DayOfTheWeek string `db:"day_of_the_week"`
  RetailerId string `db:"retailer_id"`
  RetailerName string `db:"retailer_name"`
  ProductId int `db:"product_id"`
  ProductDetailId int `db:"product_detail_id"`
  ProductUrl string `db:"product_url"`
  Brand string `db:"brand"`
  Category string `db:"category"`
  ModelNumber string `db:"model_number"`
  ProductCode string `db:"product_code"`
  SkuCode string `db:"sku_code"`
  ProductName string `db:"product_name"`
  ProductDescription string `db:"product_description"`
  Barcode string `db:"barcode"`
  Color string `db:"color"`
  Status string `db:"status"`
  Size string `db:"size"`
  Gender string `db:"gender"`
  NoOfGears string `db:"no_of_gears"`
  Suspension string `db:"suspension"`
  ImageUrlThumbnail string `db:"image_url_thumbnail"`
  ImageUrlPreview string `db:"image_url_preview"`
  Currency string `db:"currency"`
  OriginalPrice float64 `db:"original_price"`
  SalePrice float64 `db:"sale_price"`
}

type ProductInfo struct {
  TargetName string
  TargetBrand string
  Products []ProductsData
}

type ProductLev struct {
  ProductSourceString string
  ProductTargetString string
}


type PriorityQueue struct {
  HighPriority []ProductsData
  MediumPriority []ProductsData
  LowPriority []ProductsData
}

