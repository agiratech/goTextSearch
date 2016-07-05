package algorithm

import (
        "github.com/agiratech/goTextSearch/common_struct"
       )

var PriorityQueue *common_struct.PriorityQueue

func GetMatchedText(p *common_struct.ProductInfo,pl *common_struct.ProductLev) {
  GroupByPriority(p,pl)
  Levenshtein(p.TargetName)
}

