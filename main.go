// This is the Login Backend
// It should have sessions
// It should have salting for the passwords
// It should have a register part not only the login part
// It should store the information on the database

package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
)

// User holds user login information
type User struct {
	Email        string
	Username     string
	HashedPasswd string
}

// Load environment variable from .env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// HashPassword hashes a password using bcrypt and a secret pepper
func HashPassword(password string, pepper string) (string, error) {
	// Combine password + pepper before hashing
	salted := password + pepper
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(salted), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// ComparePassword compares a plain password with a hashed password
func ComparePassword(inputPassword string, pepper string, hashedPassword string) bool {
	salted := inputPassword + pepper
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(salted))
	return err == nil
}

func Thank() {
	fmt.Println("Thank you for using the system")
}

func main() {
	loadEnv()

	// Get pepper from env
	pepper := os.Getenv("PEPPER")

	// Simulate registration: pre-hash user passwords
	users := map[string]User{}

	// Pre-register 3 users with salted + hashed passwords
	rawUsers := []struct {
		email    string
		username string
		password string
	}{
		{"admin@gmail.com", "admin", "password123"},
		{"user1@gmail.com", "user1", "mypass123"},
		{"guest@gmail.com", "guest", "welcome123"},
	}

	for _, u := range rawUsers {
		hashed, err := HashPassword(u.password, pepper)
		if err != nil {
			log.Fatal("Failed to hash password:", err)
		}
		users[u.username] = User{u.email, u.username, hashed}
	}

	// Login attempt logic
	maxAttempts := 3
	attempts := 0

	for attempts < maxAttempts {
		var inputUser, inputPass string

		fmt.Println("Welcome to the Secure Login System")
		fmt.Print("Enter your username: ")
		fmt.Scanln(&inputUser)
		fmt.Print("Enter your password: ")
		fmt.Scanln(&inputPass)

		user, exists := users[inputUser]
		if !exists {
			fmt.Println("Error: Username not found.")
		} else {
			if ComparePassword(inputPass, pepper, user.HashedPasswd) {
				fmt.Println("\n✅ Login successful! Welcome,", inputUser)
				Thank()
				return
			} else {
				fmt.Println("❌ Incorrect password.")
			}
		}

		attempts++
		remaining := maxAttempts - attempts
		if remaining > 0 {
			fmt.Printf("You have %d attempt(s) remaining.\n\n", remaining)
		} else {
			fmt.Println("\n🚫 Maximum login attempts reached. System Locked.")
		}
	}
}
