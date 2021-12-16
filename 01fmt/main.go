package main

import "fmt"

func main(){
	fmt.Println("Hello Gopher")
	fmt.Printf("Hello %s!!!\n","World")
	// แบบ %v ใช้แทนได้ทุก type
	fmt.Printf("Test Number %v Test String %v ", 23 , "Age")
}