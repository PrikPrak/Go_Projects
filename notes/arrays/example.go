package main

import "fmt"

func main() {
  // Creates an array of three elements, initialized to the zero/default
  // value of whatever the specific data type is. In this case it's the
  // integer 0.
  var numberArray [3]int 

  // In Go, printing an array will print the ENTIRE array. 
  // Output of the following statement will be [0 0 0]
  fmt.Println(numberArray)

  numberArray[0] = 9
  fmt.Println(numberArray[0])


  // Attempting to access an out of bounds element will result
  // in an error.
  //numberArray[10] = 1

  
  fruitArray := [2]string{"Apples", "Oranges"}
  fmt.Println(fruitArray)

  var secondFruitArray [2]string 
  secondFruitArray = {"Bananas", "Pineapples"}
  fmt.Println(secondFruitArray)
}
