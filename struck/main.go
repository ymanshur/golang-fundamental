package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	// user := User{}
	user := User{1, "Zelda", "Safitri", "zelda@nintendo.com", true}
	// user := User{
	// 	ID:        1,
	// 	FirstName: "Zelda",
	// 	LastName:  "Safitri",
	// 	Email:     "zelda@nintendo.com",
	// 	IsActive:  true,
	// }
	// user.ID = 1
	// user.FirstName = "Zelda"
	// user.LastName = "Safitri"
	// user.Email = "zelda@nintendo.com"
	// user.IsActive = true

	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Link"
	user2.LastName = "Awekening"
	user2.Email = "link@nintendo.com"
	user2.IsActive = true

	fmt.Println(displayUser(user))
	fmt.Println(user2)
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}
