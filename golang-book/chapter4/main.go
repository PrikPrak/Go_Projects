package main

import "fmt"

func main(){
  var x string
  x = "Hello World"
  fmt.Println(x)

  x = "first"
  fmt.Println(x)

  x = "Second"
  fmt.Println(x)

  x = x + " Third"
  fmt.Println(x)

  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}