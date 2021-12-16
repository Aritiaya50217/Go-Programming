package main

import "fmt"

func main() {
	num := 15

	if num%2 == 0 {
		fmt.Println("Even")
	} else if num <= 20 {
		fmt.Println("num less than or equal to 20")
	} else {
		fmt.Println("ODD")
	}
}
