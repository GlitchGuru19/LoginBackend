// This is the Login Backend
// It should have sessions
// It should have salting for the passwords
// It should have a register part not only the login part
// It should store the information on the database

package main

import (
	"fmt"
)

// User struct defines the structure of a user in the system.
// Each user has an Email, Username, and Password.
type User struct {
	Email    string
	Username string
	Password string
}

// Message displays a welcome message when the system starts.
func Message() {
	fmt.Println("Welcome to the system")
	fmt.Println("This will contain sessions")
	fmt.Println("This will also contain registration")
}

// Thank displays a thank-you message after successful login.
func Thank() {
	fmt.Println("Thank you for using the system")
}

func main() {
	// users is a map storing User data.
	// Key: Username
	// Value: Corresponding User struct (with email, username, password)
	users := map[string]User{
		"admin": {"admin@gmail.com", "admin", "password123"},
		"user1": {"user1@gmail.com", "user1", "mypass123"},
		"guest": {"guest@gmail.com", "guest", "welcome123"},
	}

	// Maximum number of allowed login attempts
	maxAttempts := 3
	// Track how many attempts the user has made
	attempts := 0

	// Start the login loop
	for attempts < maxAttempts {
		// Variables to store user input
		var inputUser, inputPass string

		// Display login prompt
		fmt.Println("Welcome to the Login System")
		fmt.Print("Enter the username: ")
		fmt.Scanln(&inputUser)

		fmt.Print("Enter the password: ")
		fmt.Scanln(&inputPass)

		// Attempt to find the user by username (key in the map)
		user, exists := users[inputUser]

		// If user not found, display error
		if !exists {
			fmt.Println("Error: Username not found.")
		} else {
			// If user exists, check if the password matches
			if user.Password == inputPass {
				// Successful login
				fmt.Println("\nLogin successful! Welcome,", inputUser)
				Thank()
				return // Exit the program after successful login
			} else {
				// Password was incorrect
				fmt.Println("Error: Incorrect Password.")
			}
		}

		// Increment failed attempts
		attempts++
		// Calculate remaining attempts
		remaining := maxAttempts - attempts

		// Inform the user how many attempts remain or lock the system
		if remaining > 0 {
			fmt.Printf("You have %d attempts remaining\n\n", remaining)
		} else {
			fmt.Println("\nMaximum login attempts reached.")
			fmt.Println("System Locked!")
		}
	}
}
