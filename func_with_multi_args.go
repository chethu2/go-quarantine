package main

import "fmt"

func main() {
	greet("Ram", 25)
}

func greet(user string, age int) {
	fmt.Printf("Welcome %s you are %d now\n", user, age)
}
