package main

import (
  "strings"
  "fmt"

)
func main(){
  var a1, a2 string
  a1 = "tênis adidas superstar foundation"
  a2 = "tênis adidas originals superstar foundation - masculino"
  result := Levenshtein(a1,a2)
  fmt.Println(result)
}
// using levenshtein algorithm to compare two string,
// The string is matched return string
// Not matched avoid the string
func Levenshtein(a1 string, a2 string) string {
  var count int = 0
  var match_count int = 0
  var title string
  input_word := strings.Split(a1, " ")
  db_word := strings.Split(a2, " ")
  for i := range input_word {
    count += 1
    word_lenth := len(input_word[i])
    // fmt.Println(input_word[i],"=", word_lenth)
    for j := range db_word {
      db_word_lenth := len(db_word[j])
      if word_lenth == db_word_lenth{
        // fmt.Println("gggggggg")
        // compare := strings.Contains(input_word[i],db_word[j])
        if strings.Contains(input_word[i],db_word[j]){
          match_count += 1
          // fmt.Println("gggggggg")
        }
      }
      // fmt.Println(db_word[j],"=", db_word_lenth)
  }
}
  if count == match_count{
    title = "best-" + a2
  }else if match_count >= 2{
    title = "average-" + a2
  }else{
    title = "worst-" + a2
  }
return title
}