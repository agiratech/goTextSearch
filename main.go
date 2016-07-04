// program start from this package
// get the input from terminal and
// the macthed string result
package main

import(
  "runtime"
  "fmt"
  "gopkg.in/alecthomas/kingpin.v2"
  "github.com/agiratech/goTextSearch/config"
  "github.com/agiratech/goTextSearch/data_groups"
  "github.com/agiratech/goTextSearch/algorithm"
  "github.com/agiratech/goTextSearch/common_struct"
)

// get input from command line and store it in variables
var (
  brand = kingpin.Arg("brand", "Brand Name.").Required().String()
  name  = kingpin.Arg("name", "Name of product.").Required().String()
)
var Common common_struct.ProductLev
var ProductInfo *common_struct.ProductInfo
// main function
func main() {
  // use all CPU cores
  runtime.GOMAXPROCS(runtime.NumCPU())
  // Set configuration values
  config.InitConfig()
  // connect to database
  var ProductData common_struct.ProductInfo
  ProductInfo = &ProductData
  config.InitDb()
  kingpin.Parse()
  Common.ProductTargetString = "tênis adidas superstar foundation"
  Common.ProductSourceString = "tênis adidas originals superstar foundation - masculino"
  ProductInfo.TargetName = *name
  ProductInfo.TargetBrand = *brand
  data_groups.BrandClassification(ProductInfo)
  fmt.Printf("%v, %s\n%v", *brand, *name,ProductInfo)
  algorithm.GetMatchedText(ProductInfo,&Common)
  // result := algorithm.Levenshtein(Common)
  // fmt.Println(result)
}
