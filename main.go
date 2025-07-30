// This is the Login Backend
// It should have sessions
// It should have salting for the passwords
// It should have a register part not only the login part
// It should store the information on the database

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserStore struct {
	Users   []User `json:"users"`
	NextID  int    `json:"next_id"`
}

const usersFile = "users.json"

func main() {
	fmt.Println("=== LocalLink Signup System ===")
	fmt.Println()

	for {
		fmt.Println("1. Sign Up")
		fmt.Println("2. View All Users (for testing)")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			signUp()
		case "2":
			viewUsers()
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
		fmt.Println()
	}
}

func signUp() {
	fmt.Println("\n--- Sign Up ---")
	reader := bufio.NewReader(os.Stdin)

	// Get user details
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter your email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	// Validate email format (basic check)
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		fmt.Println("❌ Invalid email format!")
		return
	}

	// Check if email already exists
	userStore := loadUsers()
	for _, user := range userStore.Users {
		if user.Email == email {
			fmt.Println("❌ Email already exists!")
			return
		}
	}

	fmt.Print("Enter your password: ")
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\n❌ Error reading password!")
		return
	}
	password := string(passwordBytes)
	fmt.Println() // New line after password input

	if len(password) < 6 {
		fmt.Println("❌ Password must be at least 6 characters long!")
		return
	}

	fmt.Print("Enter your role (user/admin): ")
	role, _ := reader.ReadString('\n')
	role = strings.TrimSpace(role)

	// Validate role
	if role != "user" && role != "admin" {
		fmt.Println("❌ Role must be either 'user' or 'admin'!")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("❌ Failed to hash password!")
		return
	}

	// Create new user
	newUser := User{
		ID:       userStore.NextID,
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	// Add user to store
	userStore.Users = append(userStore.Users, newUser)
	userStore.NextID++

	// Save to file
	err = saveUsers(userStore)
	if err != nil {
		fmt.Println("❌ Failed to save user data!")
		return
	}

	fmt.Printf("✅ User '%s' created successfully with ID: %d\n", name, newUser.ID)
}

func loadUsers() UserStore {
	var userStore UserStore

	// Check if file exists
	if _, err := os.Stat(usersFile); os.IsNotExist(err) {
		// File doesn't exist, return empty store
		userStore.NextID = 1
		return userStore
	}

	// Read file
	data, err := os.ReadFile(usersFile)
	if err != nil {
		fmt.Printf("❌ Error reading users file: %v\n", err)
		userStore.NextID = 1
		return userStore
	}

	// Parse JSON
	err = json.Unmarshal(data, &userStore)
	if err != nil {
		fmt.Printf("❌ Error parsing users file: %v\n", err)
		userStore.NextID = 1
		return userStore
	}

	// Ensure NextID is set properly
	if userStore.NextID == 0 {
		userStore.NextID = len(userStore.Users) + 1
	}

	return userStore
}

func saveUsers(userStore UserStore) error {
	data, err := json.MarshalIndent(userStore, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(usersFile, data, 0644)
}

func viewUsers() {
	fmt.Println("\n--- All Users ---")
	userStore := loadUsers()

	if len(userStore.Users) == 0 {
		fmt.Println("No users found.")
		return
	}

	fmt.Printf("%-5s %-20s %-30s %-10s\n", "ID", "Name", "Email", "Role")
	fmt.Println(strings.Repeat("-", 70))

	for _, user := range userStore.Users {
		fmt.Printf("%-5d %-20s %-30s %-10s\n", user.ID, user.Name, user.Email, user.Role)
	}
}