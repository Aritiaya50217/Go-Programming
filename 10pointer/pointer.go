package main

import "fmt"

func main() {
	me := "Gopher"
	fmt.Printf("You are %s\n", me)

	// ใช้ & เพื่อแสดงที่อยู่บน memory
	addr := &me
	fmt.Println(addr)
	fmt.Printf("%T\n", addr)

	// แทนที่ข้อมูลลงใน memory
	// กรณีแทนที่ลงใน memory ที่เดียวกับ me
	*addr = "Penguin"
	fmt.Printf(me)
}
