package data_groups

import (
  "fmt"
  "github.com/agiratech/goTextSearch/config"
)

func BrandClassification(brand string){
  fmt.Println("brand classification", brand, config.Db)
}