package main

import (
	"fmt"
)

// Define a User struct to hold individual user data
type User struct {
	Email    string
	Username string
	Password string
}

// Welcome message function
func Message() {
	fmt.Println("Welcome to the system")
	fmt.Println("This will contain sessions")
	fmt.Println("This will also contain registration")
}

// Thank you message
func Thank() {
	fmt.Println("Thank you for using the system")
}

// Main program
func main() {
	// Slice of Users
	users := []User{
		{"admin@gmail.com", "admin", "password123"},
		{"user1@gmail.com", "user1", "mypass123"},
		{"guest@gmail.com", "guest", "welcome123"},
	}

	maxAttempts := 3
	attempts := 0

	for attempts < maxAttempts {
		var inputUser, inputPass string

		fmt.Println("Welcome to the Login System")
		fmt.Println()
		fmt.Print("Enter the username: ")
		fmt.Scanln(&inputUser)

		fmt.Print("Enter the password: ")
		fmt.Scanln(&inputPass)

		// Check login
		userFound := false

		for _, user := range users {
			if user.Username == inputUser {
				userFound = true
				if user.Password == inputPass {
					fmt.Println("\nLogin successful! Welcome,", inputUser)
					Thank()
					return
				} else {
					fmt.Println("Error: Incorrect Password.")
					break
				}
			}
		}

		if !userFound {
			fmt.Println("Error: Username not found.")
		}

		attempts++
		remaining := maxAttempts - attempts
		if remaining > 0 {
			fmt.Printf("You have %d attempts remaining\n\n", remaining)
		} else {
			fmt.Println("\nMaximum login attempts reached.")
			fmt.Println("System Locked!")
		}
	}
}
