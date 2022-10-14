package main

import (
	"fmt"
	"struck/management"
)

func main() {
	// user := User{}
	// user.ID = 1
	// user.FirstName = "Zelda"
	// user.LastName = "Safitri"
	// user.Email = "zelda@nintendo.com"
	// user.IsActive = true
	// user := User{
	// 	ID:        1,
	// 	FirstName: "Zelda",
	// 	LastName:  "Safitri",
	// 	Email:     "zelda@nintendo.com",
	// 	IsActive:  true,
	// }
	// user := User{1, "Zelda", "Safitri", "zelda@nintendo.com", true}
	user := management.User{1, "Zelda", "Safitri", "zelda@nintendo.com", true}
	user2 := management.User{2, "Link", "Awekening", "link@nintendo.com", true}

	// fmt.Println(user)
	// fmt.Println(displayUser(user))
	// fmt.Println(user2)
	fmt.Println(user2.Display())

	users := []management.User{user, user2}

	group := management.Group{"Gamer", user, users, true}

	// displayGroup(group)
	group.Display()
}

// func displayUser(user User) string {
// 	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
// }

// func displayGroup(group Group) {
// 	fmt.Printf("Name : %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count : %d", len(group.Users))
// 	fmt.Println("")

// 	fmt.Println("Users name :")
// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}
// }
