/*
Problem #5: The fscanf Function
Author: Corey Prak
Date Created: 04/03/2014
Comments:
Had to refer to documentation so much. Really not used to the organization:
func Open(name string) (file *File, err error)

This is how I see it, hope it helps:
func *function name*(*argument name and data type*) (*return parameters, name then datatype*)
*/

package main

//importing fmt for file I/O and os for opening file. 
import(
fmt
os
)


func main(){

  // Open, used for reading,
  //takes a string, returns a pointer to a file (file) & error (err)
  filePtr := os.Open("testdata6")

  

  
}