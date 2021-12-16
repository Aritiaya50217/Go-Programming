package main

import "fmt"

// struct ชุดของข้อมูล
type User struct {
	Username      string
	FullName      string
	ProfilePicUrl string
}

func main() {
	username := "golang"
	fullName := "Basic Golang"
	profilePicUrl := "https://knowhere.fack/gopher.jpg"
	fmt.Println(username, fullName, profilePicUrl)

	// u := User{}
	// u.Username = username
	// u.FullName = fullName
	// u.ProfilePicUrl = profilePicUrl

	u := User{
		Username:      username,
		FullName:      fullName,
		ProfilePicUrl: profilePicUrl,
	}

	fmt.Printf("% #v\n", u)

	// ดูเฉพาะค่าที่ต้องการ
	name := u.Username
	fmt.Println(name)
	fmt.Println(u.FullName)

}
