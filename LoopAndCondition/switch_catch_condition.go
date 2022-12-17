package main

import "fmt"

func main() {

	name := "2"
	switch name {
	case "1":
		fmt.Println("one")
	case "2":
		fmt.Println("two")
	default:
		fmt.Println("nothing")
	}
}
