package algorithm

import (
  "strings"
  _"fmt"
  "github.com/agiratech/goTextSearch/common_struct"
)

// using levenshtein algorithm to compare two string,
// The string is matched return string
// Not matched avoid the string
var recommended_product common_struct.PriorityQueue

func Levenshtein(target_str string) {
  var source_word,target_word string
  var source_words,matched []string
  var cost int
  var percentage float64
  target_words := strings.Split(target_str, " ")
  for _,product = range PriorityQueue.HighPriority{
    source_words = strings.Split(product.ProductName, " ")
    for _,target_word = range target_words {
      for _,source_word = range source_words {
        cost = LevenshteinDistance(target_word,source_word)
        percentage = GetPercentage(cost,len(target_word))
        if percentage <= 25.0 {
          matched = append(matched,source_word)
        }
      }
    }
    percentage = GetPercentage(len(matched),len(target_words))
    GetResult(product,percentage)
    matched = nil
  }
}

func min(a1 int, a2 int, a3 int) int {
  min := a1
  if a2 < min {
    min = a2
  }
  if a3 < min {
    min = a3
  }
  return min
}

func make2d(l1 int, l2 int) [][]int {
  ret := make([][]int, l1)
  for i := 0; i < l1; i++ {
    ret[i] = make([]int, l2)
  }
  return ret
}

func LevenshteinDistance(str1 string, str2 string) int {
  distance := make2d(len(str1)+1, len(str2)+1)
  for i := 0; i <= len(str1); i++ {
    distance[i][0] = i;
  }
  for i := 0; i <= len(str2); i++ {
    distance[0][i] = i;
  }

  for i := 1; i <= len(str1); i++ {
    for j := 1; j <= len(str2); j++ {
      ex := 1
      if str1[i - 1] == str2[j - 1] {
        ex = 0
      }
      distance[i][j] = min(distance[i - 1][j] + 1,
                           distance[i][j - 1] + 1,
                                 distance[i - 1][j - 1] + ex)
    }
  }

  return distance[len(str1)][len(str2)]

}

func GetPercentage(cost int,len int) float64{
  return ((float64(cost) / float64(len))*100)
}

func GetResult(product common_struct.ProductsData, percentage float64) {
  switch {
    case percentage >= 75: recommended_product.HighPriority = append(recommended_product.HighPriority,product)
                           break;
    case (percentage < 75 && percentage >= 25 ): recommended_product.MediumPriority = append(recommended_product.MediumPriority,product)
                           break;
    case percentage < 25: recommended_product.LowPriority = append(recommended_product.LowPriority,product)
                           break;

  }
  RecommendedProducts = &recommended_product

}