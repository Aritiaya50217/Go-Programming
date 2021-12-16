package main

import "fmt"

func main() {
	const before = "Before"
	fmt.Println("Status : ", before)

	// const ไม่สามารถเปลี่ยนแปลงค่าได้
	// before = "After"
	// fmt.Println(before)

	// iota คือ การ assign ค่าเป็นลำดับ ลำดับถัดไป เช่น 0 1 2 3

	const (
		One = iota
		Two
		Three
	)

	fmt.Println(One)
	fmt.Println(Two)
	fmt.Println(Three)

}
