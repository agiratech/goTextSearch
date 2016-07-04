package algorithm

import (
        "fmt"
        "github.com/agiratech/goTextSearch/common_struct"
        "strings"
       )

type TextDetail struct {
  TargentNameLen int
  SourceNameLen int
  TargetWords []string
  SourceWords []string
  MatchedStr []string
}

var target_word,source_word string
var product common_struct.ProductsData
var text_detail TextDetail
func GroupByPriority(p *common_struct.ProductInfo,pl *common_struct.ProductLev) {
  text_detail.TargetWords= strings.Split(pl.ProductTargetString," ")
  for _,product = range p.Products {
    text_detail.SourceWords = strings.Split(product.ProductName," ")
    text_detail.TargentNameLen = 0
    text_detail.SourceNameLen = 0
    for _,target_word = range text_detail.TargetWords {
      text_detail.TargentNameLen = len(target_word)
      for _,source_word = range text_detail.SourceWords {
        text_detail.SourceNameLen = len(source_word)
        if text_detail.SourceNameLen == text_detail.TargentNameLen && string(target_word[0]) == string(source_word[0]){
          text_detail.MatchedStr = append(text_detail.MatchedStr,source_word)
        }
      }
    }
    // ((len(text_detail.MatchedStr) /text_detail.TargentNameLen)*100)
    fmt.Println(len(text_detail.MatchedStr),text_detail.TargentNameLen)
    // text_detail.GetPercentage()
    text_detail.MatchedStr = nil
    text_detail.SourceWords = nil
    text_detail.MatchedStr = nil
  }
}
