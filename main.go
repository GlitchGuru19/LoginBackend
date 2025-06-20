package main

import (
	"fmt"
	// packages for salting and hashing the passwords,
	// creatig access tokens and connecting to a database
)

func main(){
	var
	(
		sUsername string
		username string = "Carlos"
		//sPassword string
		password string = "Falcon12@"
	)
	fmt.Println("\nWelcome to the Login System")
	fmt.Println("\n\nPlease Enter your credentials")
	fmt.Print("\nEnter your username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	//fmt.Printf("This is your username: %s \nand this is your password: %s", username, password)
	if username == sUsername {
		fmt.Println("Access granted")
	}
}