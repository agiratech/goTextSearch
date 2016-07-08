package algorithm

import (
        "github.com/agiratech/goTextSearch/common_struct"
        "fmt"
        "encoding/json"

         )

type TextDetail struct {
  TargentNameLen int
  SourceNameLen int
  TargetWords []string
  SourceWords []string
  MatchedStr []string
}

var RecommendedProducts,PriorityQueue *common_struct.PriorityQueue
var Recommend_product *common_struct.PriorityQueue
func GetMatchedText(p *common_struct.ProductInfo,pl *common_struct.ProductLev) (*common_struct.PriorityQueue) {
  GroupByPriority(p,pl)
  fmt.Println(PriorityQueue.LowPriority,PriorityQueue.MediumPriority)
  Levenshtein(pl.ProductTargetString)

  recommend,_ := json.Marshal(RecommendedProducts)
  if string(recommend) !="null" {
    Recommend_product = &common_struct.PriorityQueue {HighPriority: RecommendedProducts.HighPriority ,MediumPriority: RecommendedProducts.MediumPriority, LowPriority: RecommendedProducts.LowPriority}
  }
  return Recommend_product

}

