package main

import "fmt"

func main() {
	langs := []string{}
	fmt.Printf("langs : %v\n", langs)

	// check slice
	if langs == nil {
		fmt.Println("Yes nil 'langs' is a nil")
	} else {
		fmt.Printf("langs in not nil , value : %v\n", langs)
	}
}
