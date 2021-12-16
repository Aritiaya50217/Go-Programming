package main

import (
	"fmt"
	"igapp/time"
	"igapp/user"
)

// package user
// func info(name string, msg string, age int) {
// 	fmt.Printf("My name is %s\n", name)
// 	fmt.Printf("My message is %s\n", msg)
// 	fmt.Printf("I'm %d year old\n", age)
// }

// package time
// func today() string {
// 	return "Today"
// }

func main() {
	user.Info("Oil", "Hello", 24)
	t := time.Today()
	fmt.Println(t)
}
