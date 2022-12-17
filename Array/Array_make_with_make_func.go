package main

import "fmt"

func main() {
	//array make with make() function
	slc := make([]int, 3, 5)
	fmt.Println(slc)
	slc[0] = 1
	fmt.Println(slc)
}
