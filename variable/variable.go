package main

import "fmt"

func main() {
	/* var x int
	x = 12
	fmt.Println(x) */
	var x float32
	x = 12.3
	fmt.Println(x)

	// variable declaration shorthand but you not set a type variable shorthand automatically known variable type (:=)
	f := "apple"
	fmt.Println(f)

	b := 1
	fmt.Println(b)

	//constant
	const xx string = "constant"
	// error
	//xx = "suvo"
	fmt.Println(xx)

	const n = 500000000
	const d = 6e+11
	fmt.Println(int64(d))
	fmt.Println(n)
}
