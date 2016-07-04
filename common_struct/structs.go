package common_struct

type Product struct {
  ProductId int
  ProductName string
  ProductPrice float64

}

type ProductInfo struct {
  TargetName string
  TargetBrand string
  Products []Product
}