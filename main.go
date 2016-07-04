// program start from this package
// get the input from terminal and
// the macthed string result
package main

import( "fmt"
        "gopkg.in/alecthomas/kingpin.v2"
        _"github.com/agiratech/goTextSearch/algorithm"
      )

// get input from command line and store it in variables
var (
  brand = kingpin.Arg("brand", "Brand Name.").Required().String()
  name  = kingpin.Arg("name", "Name of product.").Required().String()
)


func main() {
  kingpin.Parse()
  fmt.Printf("%v, %s\n", *brand, *name)
}
