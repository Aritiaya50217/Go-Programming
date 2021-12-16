package main

import "fmt"

func showAll(sa [4]string) {
	fmt.Printf("Show all : %v\n", sa)
}

func main() {
	// array
	var langs = [4]string{"golang", "python", "javascript"}
	fmt.Printf("Langs : %v\n", langs)

	// เข้าถึงลำดับใน array
	n := langs[2]
	fmt.Println("Index 2 :", n)

	// เปลี่ยนแปลงค่า
	langs[1] = "Java"
	fmt.Printf("After: %v\n", langs)

	// ต้องมีขนาด array เท่ากันจึงจะเรียกใช้ฟังก์ชันที่มีอาร์เรย์ได้
	cords := [4]string{"a", "b", "c"}
	showAll(cords)
}
