package main

import "fmt"

func info(name, msg string) {
	fmt.Printf("My name is %s %s\n", name, msg)
}

func today() string {
	return "Today"
}

func swap(x, y string) (string, string) {
	return x, y
}

func main() {
	// call function
	info("Gopher", "!!!")
	info("Oil", "???")

	day := today()
	fmt.Println(day)

	a, b := swap("Go", "pher")
	fmt.Println(a, b)
}
