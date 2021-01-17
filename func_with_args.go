package main

import "fmt"

func main() {
	greet("Ram")
}

func greet(user string) {
	fmt.Printf("Welcome %s\n", user)
}
