/*
Problem #20: Area of a Circle 
Author: Corey Prak
Date Created: 04/21/2014
Comments:
  Still need to figure out the %q format specifier...
*/

package main

import(
"fmt"
"math"
)

func main() {
  var radius float32
  var scanErr error

  fmt.Printf("\nPlease enter a floating point value for the radius of a circle: ")
  _, scanErr = fmt.Scanf("%f", &radius)
  if scanErr != nil {
    fmt.Printf("\nERROR: %q.\n\nExiting...\n\n", scanErr)

    return
  }

  fmt.Printf("\nThe area of a circle with radius %f is %f.\n\n", radius, ((radius * 2) * math.Pi))  
}
