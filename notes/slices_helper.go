package main

import "fmt"

func main(){
  
  numArray := [3]int {2, 3, 4}

  var exampleSlice []int 
  exampleSlice = numArray[:]

  fmt.Printf("\n\nnumArray of type %T.\n\n", numArray)

  fmt.Printf("\n\nexampleSlice of type %T.\n\n", exampleSlice)
}
