// program start from this package
// get the input from terminal and
// the macthed string result
package main

import(
  "runtime"
  // "fmt"
  "github.com/agiratech/goTextSearch/config"
  "github.com/agiratech/goTextSearch/data_groups"
  "github.com/agiratech/goTextSearch/algorithm"
  "github.com/agiratech/goTextSearch/common_struct"
  "strings"
  "github.com/gin-gonic/gin"
  "github.com/itsjamie/gin-cors"
  "time"
)

var Common common_struct.ProductLev
var ProductInfo *common_struct.ProductInfo
var ReturnValue *common_struct.PriorityQueue
var ProductLevName string

// main function
func main() {
  // use all CPU cores
  runtime.GOMAXPROCS(runtime.NumCPU())
  // Set configuration values
  config.InitConfig()
  // connect to database
  config.InitDb()
  route := gin.Default()
  route.Use(cors.Middleware(cors.Config{
    Origins:        "*",
    Methods:        "GET, PUT, POST, DELETE",
    RequestHeaders: "Origin, Authorization, Content-Type",
    ExposedHeaders: "",
    MaxAge: 50 * time.Second,
    Credentials: true,
    ValidateHeaders: false,
  }))
  route.GET("/v2",handler)
  route.Run(":8001")
}

func handler(server *gin.Context) {
  brand := server.Query("brand")
  title := server.Query("title")
  color := server.Query("color")
  ReturnValue = GetProductInfo(brand, color, title)
  server.JSON(200, gin.H{"priorities": ReturnValue})
}

func GetProductInfo(brand string, color string, title string) *common_struct.PriorityQueue {
  var ProductData common_struct.ProductInfo
  ProductInfo = &ProductData
  Common.ProductTargetString = strings.ToLower(title)
  ProductInfo.TargetName = title
  ProductInfo.TargetBrand = strings.ToLower(brand)
  ProductInfo.TargetColor = strings.ToLower(color)
  data_groups.BrandClassification(ProductInfo)
  ReturnValue = algorithm.GetMatchedText(ProductInfo,&Common)
  return ReturnValue
}
