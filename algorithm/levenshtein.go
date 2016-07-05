package algorithm

import (
  "strings"
  // "fmt"
  _"github.com/agiratech/goTextSearch/common_struct"
)


// func main(){
//   var Common common_struct.ProductLev
//   // var a1, a2 string
//   Common.ProductTargetString = "tênis adidas superstar foundation"
//   Common.ProductSourceString = "tênis adidas originals superstar foundation - masculino"
//   result := Levenshtein(Common)
//   fmt.Println(result)
// }
// using levenshtein algorithm to compare two string,
// The string is matched return string
// Not matched avoid the string
func Levenshtein(str string) {
  // var re_str common_struct.ProductLevString
  var count int = 0
  var match_count int = 0
  input_word := strings.Split(str, " ")
  for _,product = range PriorityQueue.HighPriority{
    db_word := strings.Split(product.ProductName, " ")
    for i := range input_word {
      count += 1
      // word_lenth := len(input_word[i])
      // fmt.Println(input_word[i],"=", word_lenth)
      for j := range db_word {
        // db_word_lenth := len(db_word[j])
        if strings.Contains(input_word[i],db_word[j]){
          match_count += 1
        }
      }
    }
  }
//   if count == match_count{
//     re_str.ProductLevName = "best-" + str
//   }else if match_count >= 2{
//     re_str.ProductLevName = "average-" + str.ProductSourceString
//   }else{
//     re_str.ProductLevName = "worst-" + str.ProductSourceString
//   }
// return re_str
}