package main

import (
	"fmt"
	"log"
)

func Readfile(name string) (string, error) {
	// read file ...
	// err := errors.New("file not found")
	// return "", err
	return "data...", nil
}

func main() {
	data, err := Readfile("profile.txt")
	if err != nil {
		// ใช้ log เพื่อแสดงเวลา
		log.Println(err)
		return
	}
	fmt.Println("read file success :", data)
}
