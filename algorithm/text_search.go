package algorithm

import (
        "github.com/agiratech/goTextSearch/common_struct"
       )

func GetMatchedText(p *common_struct.ProductInfo,pl *common_struct.ProductLev) {
  GroupByPriority(p,pl)
}

