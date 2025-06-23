package main

import (
	"fmt"
	// packages for salting and hashing the passwords,
	// creatig access tokens and connecting to a database
)

func Message(){
	fmt.Println("Welcome to the system")
	fmt.Println("This will contain sessions")
	fmt.Println("This will also contain registration")
}

func main(){

	// User credentials stored in parallel slices
	usernames := []string{"admin", "user1", "guest"}
	passwords := []string{"password123","mypass123","welcome123"}

	// Login attempts
	maxAttempts := 3
	attempts := 0

	for attempts < maxAttempts {
		// Get user input
		var (
			inputUser string
			inputPass string
		)

		fmt.Println("Welcome to the Login System")
		fmt.Println()
		fmt.Print("Enter the username: ")
		fmt.Scanln(&inputUser)

		fmt.Print("Enter the password: ")
		fmt.Scanln(&inputPass)

		// Check credentials 
		userFound := false
		correctPassword := false
		var userIndex int

		// Search for username
		for i, username := range usernames{
			if username == inputUser{
				userFound = true
				userIndex = i
				break
			}
		}

		// Check password if user is found
		if userFound {
			if passwords[userIndex] == inputPass {
				fmt.Println("\nLogin succesful! Welcome, ", inputUser)
				return
			} else {
				correctPassword = false
			}
		}

		// Give appropriate error message
		switch {
		case !userFound:
			fmt.Println("Error: Username not found.")
		case !correctPassword:
			fmt.Println("Error: Incorrect Password.")
		}
		attempts++
		remaining := maxAttempts - attempts
		if remaining > 0 {
			fmt.Printf("You have %d attempts remaining\n", remaining)
		} else {
			fmt.Println("\nMaximum login attempts reached")
		}
	}
}