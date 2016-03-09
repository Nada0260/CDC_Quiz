//Package coder is to be written
package coder

import "fmt"

//z is a public global variable

var z int

//Sum is a function
func Sum(x, y int) int {
	z = x + y
	fmt.Println(z)
	return z
}
