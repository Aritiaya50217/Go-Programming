package main

import "fmt"

var name string = "Gopher !!!!"

func main() {
	// ตัวแปร := ใช้ได้แค่ใน bodyเท่านั้น
	num := 123
	fmt.Println(name)
	fmt.Printf("String : %s\n", name)
	fmt.Printf("Type : %T", num)

	/* basic type
	bool
	string

	int int8 int16 int32 int64
	unit unit8 uint16 unit32 unit64 unitptr

	byte // alias for unit8

	rune // alias for int32
		//represents a Unicode code point

	float32 float64

	complex64 complex128
	*/

	i := 12
	s := "String"
	f := 2.22
	b := true
	r := "x"

	fmt.Printf("Int : %d\n", i)
	fmt.Printf("String : %s\n", s)
	fmt.Printf("float : %f\n", f)
	fmt.Printf("bool : %v\n", b)
	fmt.Printf("rune : %v", r)
}
