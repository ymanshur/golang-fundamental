package main

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) Display() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

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
	user := User{1, "Zelda", "Safitri", "zelda@nintendo.com", true}
	user2 := User{2, "Link", "Awekening", "link@nintendo.com", true}

	// fmt.Println(user)
	fmt.Println(displayUser(user))
	// fmt.Println(user2)
	fmt.Println(user2.Display())

	users := []User{user, user2}

	group := Group{"Gamer", user, users, true}

	// displayGroup(group)
	group.Display()
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

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
