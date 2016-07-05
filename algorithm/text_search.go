package algorithm

import (
        "github.com/agiratech/goTextSearch/common_struct"
        "fmt"
       )

type TextDetail struct {
  TargentNameLen int
  SourceNameLen int
  TargetWords []string
  SourceWords []string
  MatchedStr []string
}

var RecommendedProducts,PriorityQueue *common_struct.PriorityQueue

func GetMatchedText(p *common_struct.ProductInfo,pl *common_struct.ProductLev) {
  GroupByPriority(p,pl)
  Levenshtein(pl.ProductTargetString)
  fmt.Println(len(RecommendedProducts.HighPriority), len(RecommendedProducts.MediumPriority),len(RecommendedProducts.LowPriority))

  for _,a := range RecommendedProducts.HighPriority {
    fmt.Println(a.ProductName,"  High")
  }
  for _,b := range RecommendedProducts.MediumPriority {
    fmt.Println(b.ProductName,"  Medium")
  }
  for _,c := range RecommendedProducts.LowPriority {
    fmt.Println(c.ProductName,"  Low")
  }
}

