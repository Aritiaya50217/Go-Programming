package main

import "fmt"

func printSlice(ns []string) {
	fmt.Printf("Print : %v\n", ns)
}

func main() {
	// slice ไม่ต้องกำหนดขนาด
	langs := []string{"golang", "python", "javascript"}
	fmt.Printf("langs :%#v\n", langs)

	s := langs[0]
	fmt.Printf("lang[0]: %s\n", s)

	// หาขนาด slice
	l := len(langs)
	fmt.Printf("Lenght : %v\n", l)

	// เพิ่มข้อมูลต่อท้าย
	langs = append(langs, "Java", "PHP")
	fmt.Printf("After : %v\n", langs)
	fmt.Println("Lenght : ", len(langs))

	// slicing
	a := langs[0:2]
	fmt.Printf("langs :%v\n", langs)
	fmt.Printf("a : %v\n", a)

	b := langs[1:3]
	fmt.Printf("b : %v\n", b)

	// ลำดับแรก ถึง ลำดับสุดท้าย
	n := langs[0:]
	fmt.Printf("n : %v\n", n)

	u := langs[:]
	fmt.Printf("u :%v\n", u)

	printSlice(langs)

	cord := [4]string{"A", "B", "C", "D"}
	// แปลงอาร์เรย์เป็น slice โดยใช้ [:]
	printSlice(cord[:])
}
