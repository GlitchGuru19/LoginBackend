package main

import (
	"fmt"
	// packages for salting and hashing the passwords,
	// creatig access tokens and connecting to a database
)

func main(){

	// stored passwords
	correctUsername := "Carlos"
	correctPassword := "Falcon12@"
	
	// Get user input
	var (
		username string
		password string
	)

	fmt.Println("Welcome to the Login System")
	fmt.Println()
	fmt.Print("Enter the username: ")
	fmt.Scanln(&username)

	fmt.Print("Enter the password: ")
	fmt.Scanln(&password)

	// Check credentials
	if username == correctUsername && password == correctPassword{
		fmt.Println("Login successful!")
	} else if username == correctUsername && password != correctPassword{
		fmt.Println("Wrong password.")
	}else if username != correctUsername && password == correctPassword{
		fmt.Println("No such username.")
	} else {
		fmt.Println("Both username and password are incorect.")
	}
}