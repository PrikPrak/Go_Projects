
/* Package declaration. Every Go program MUST start with one.
package main is Go's way of organizing and reusing code. There are two types of 
Go programs: executables and libraries. */
package main

/* Include code from other packages. "fmt" package is for
formatting input and output. Anything in double quotes is a string literal.
import "fmt". Enter command godoc fmt Println for more info. */

// Function delcaration
func main(){
  fmt.Println("Hello World")
}