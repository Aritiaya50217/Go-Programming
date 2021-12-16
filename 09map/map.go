package main

import "fmt"

func main() {
	// map คล้ายกับ dictionary
	// กรณีนี้ map มี key เป็น int ส่วน value เป็น string
	status := map[int]string{
		200: "OK",
		300: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
		600: "",
	}
	fmt.Printf("% #v\n", status)

	// นับขนาดของ map
	l := len(status)
	fmt.Printf("lenght : % #v\n", l)

	// แก้ไขข้อมูลใน map
	status[200] = "Okey"
	fmt.Printf("% #v\n", status)

	// add
	status[211] = "Add Status"
	fmt.Printf("New Status :% #v\n", status)

	// อ่านแค่ value
	value := status[211]
	fmt.Printf("Value:%#v\n", value)

	// delete
	delete(status, 211)
	fmt.Printf("% #v\n", status)

	// check value
	if v, ok := status[600]; ok {
		fmt.Printf("value:%q\n\n", v)
	} else {
		fmt.Println("not found")
	}

	// allocates ก่อนใช้งาน โดยใช้ make หรือจะ กำหนดค่าลงไปเลยก็ได้
	// var m = make(map[string]string)
	m := map[string]string{
		"Test": "ทดสอบ",
	}
	if m == nil {
		fmt.Printf("m is nil , value:% #v\n", m)
	} else {
		fmt.Printf("m is not nil. value : %#v\n", m)
	}
}
