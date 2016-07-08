package algorithm

import (
        _"fmt"
        "github.com/agiratech/goTextSearch/common_struct"
        "strings"
       )

var target_word,source_word string
var product common_struct.ProductsData
var priority_queue common_struct.PriorityQueue
var percentage float64
func GroupByPriority(p *common_struct.ProductInfo,pl *common_struct.ProductLev) {
  text_detail := &TextDetail{TargetWords:strings.Split(pl.ProductTargetString," ")}
  for _,product = range p.Products {
    text_detail.SourceWords = strings.Split(product.ProductName," ")
    text_detail.TargentNameLen = 0
    text_detail.SourceNameLen = 0
    for _,target_word = range text_detail.TargetWords {
      text_detail.TargentNameLen = len(target_word)
      for _,source_word = range text_detail.SourceWords {
        text_detail.SourceNameLen = len(source_word)
        if text_detail.SourceNameLen > 0 && text_detail.TargentNameLen > 0 && string(target_word[0]) == string(source_word[0]){
          text_detail.MatchedStr = append(text_detail.MatchedStr,source_word)
        }
      }
    }
  text_detail.GetPercentage(product)
  }
}


func (text_detail *TextDetail) GetPercentage(product common_struct.ProductsData) {
  percentage = ((float64(len(text_detail.MatchedStr))) / float64(len(text_detail.TargetWords)))*100
  switch {
    case percentage >= 60: priority_queue.HighPriority = append(priority_queue.HighPriority,product)
                           break;
    case (percentage < 60 && percentage >= 25 ): priority_queue.MediumPriority = append(priority_queue.MediumPriority,product)
                           break;
    case percentage < 25: priority_queue.LowPriority = append(priority_queue.LowPriority,product)
                           break;

  }
  PriorityQueue = &priority_queue
  text_detail.MatchedStr = nil
  text_detail.SourceWords = nil
  text_detail.MatchedStr = nil
}