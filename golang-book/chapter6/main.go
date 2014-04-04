package main

import "fmt"

func main(){
  var x [5]int
  x[4] = 100
  fmt.Println(x)

  var y [5]float64
  y[0] = 98
  y[1] = 93
  y[2] = 77
  y[3] = 82
  y[4] = 83
 
  var total float64 = 0

/*
  for i := 0; i < len(y); i++ {
    total += y[i]
  }
*/

/*i - current position in array
  value - the same as x[i] */
  for _, value := range y {
    total += value
  }

  fmt.Println(total / float64(len(y)))

  test_map := make(map[string]int)
  test_map["key"] = 10
  fmt.Println(test_map)
  fmt.Println(test_map["key"])
}

