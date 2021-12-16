package main

import "fmt"

// struct ชุดของข้อมูล
type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

// info() = method
func (u User) info() {
	username := "golang"
	fullName := "Basic Golang"
	profilePicUrl := "https://knowhere.fack/gopher.jpg"
	fmt.Println(username, fullName, profilePicUrl)
}

func main() {
	u := User{
		Username:      "golang",
		FullName:      "Basic Golang",
		ProfilePicUrl: "https://knowhere.fack/gopher.jpg",
	}

	u.info()
}
