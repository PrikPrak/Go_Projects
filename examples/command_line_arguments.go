package main

import(
  "fmt"
  "os"
)

func main() {
  var i int

  fmt.Println("Of the Args variable of the os package which is of type string slice, there are/is", len(os.Args), "argument(s)")

  for i = 0; i < len(os.Args); i++ {
    fmt.Println("Args[", i, "] = ", os.Args[i])
  }
}
