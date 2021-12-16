package main

import "fmt"

func main() {
	// forloop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// กรณี while loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	langs := [5]string{"golang", "python", "javascript"}
	fmt.Println("\nfor basic")

	for i := 0; i < len(langs); i++ {
		value := langs[i]
		fmt.Println(i, ":", value)
	}

	fmt.Println("\nfor range slice")
	// ใช้ range จะ return index and value
	for index, value := range langs {
		fmt.Println(index, ":", value)
	}

	// กรณีไม่อยากแสดง index
	fmt.Println("\nonly value")
	for _, value := range langs {
		fmt.Println("value:", value)
	}

	// กรณีไม่อยากแสดง value
	fmt.Println("\nonly index")
	for index := range langs {
		fmt.Println("index:", index)
	}
}
