package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// function value
var subtraction = func(a int, b int) int { return a - b }

// กำหนด type แบบ function ตามเงื่อนไข
func cal(op func(int, int) int) {
	r := op(4, 5)
	fmt.Println(r)
}

func main() {
	var name string = "Oil"
	fmt.Printf("value: %v\n", name)
	fmt.Printf("Type: %T\n", name)

	p := plus(1, 2)
	fmt.Println("1 + 2 :", p)

	sub := subtraction(5, 4)
	fmt.Println("a - b =", sub)
	fmt.Printf("Type :%T\n", sub)

	test := func(a int, b int) int { return a * b }
	t := test(1, 2)
	fmt.Printf("Result : %d", t)

	minus := func(a, b int) int { return a - b }

	// function cal เรียกใช้ function อื่น
	cal(test)
	cal(minus)
}
