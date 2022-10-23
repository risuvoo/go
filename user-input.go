package main

import "fmt"

func main() {
	var value int
	fmt.Scan(&value)
	fmt.Println("given value is ", value)

	fmt.Scanln(&value)
	fmt.Println("another value is ", value)

	fmt.Scanf("%d", &value) //formatting like c program
	fmt.Println("formating value is ", value)
}
