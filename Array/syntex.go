package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 0}
	fmt.Println(arr)
	//dynamic len
	const len = 2
	arr2 := [len]string{"suvo", "moin"}
	fmt.Println(arr2)
	//empty initialization
	arr3 := [5]int{}
	fmt.Println(arr3)

	arr4 := [5]string{}
	fmt.Println(arr4)
}
