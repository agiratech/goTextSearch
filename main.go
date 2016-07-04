// program start from this package
// get the input from terminal and
// the macthed string result
package main

import(
  "runtime"
  "fmt"
  "gopkg.in/alecthomas/kingpin.v2"
  "github.com/agiratech/goTextSearch/config"
  "github.com/agiratech/goTextSearch/data_groups"
  _"github.com/agiratech/goTextSearch/algorithm"
)

// get input from command line and store it in variables
var (
  brand = kingpin.Arg("brand", "Brand Name.").Required().String()
  name  = kingpin.Arg("name", "Name of product.").Required().String()
)


func main() {
  // use all CPU cores
  runtime.GOMAXPROCS(runtime.NumCPU())

  // Set configuration values
  config.InitConfig()

  // connect to database
  data_groups.InitDb()

  kingpin.Parse()
  data_groups.BrandClassification(*brand)
  fmt.Printf("%v, %s\n", *brand, *name)
}
