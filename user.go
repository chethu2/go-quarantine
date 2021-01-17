package main

import (
	"fmt"
)

func main() {
	var userName, password string
	fmt.Println("Please enter user name")
	fmt.Scanf("%s", &userName)
	fmt.Println("Please enter password")
	fmt.Scanf("%s", &password)

	if len(userName) < 3 {
		fmt.Printf("User name %s is not valid: user name should be at least more than 3 characters.\n", userName)
		return
	}
	if len(password) < 8 {
		fmt.Printf("Password %s is not valid: password should be at least more than 8 characters.\n", password)
		return
	}
	fmt.Println("Congratulations!! User name and password are valid!")
}
