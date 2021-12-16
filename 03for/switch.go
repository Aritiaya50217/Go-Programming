package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday
	// เลือกให้ตรงกับเหตุการณ์
	switch today {
	case time.Monday:
		fmt.Println("today is Monday")
		// fallthrough ทำเงื่อนไขด้านล่างด้วย กรณีเข้าเงื่อนไขด้านบน
		fallthrough
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Printf("today is %v", today)
	}
}
