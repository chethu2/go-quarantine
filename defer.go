package main

import "fmt"

func main() {
	defer openDoor()
	defer closeDoor()
	enterHouse()
}

func openDoor() {
	fmt.Println("Opening the Door")
}

func enterHouse() {
	fmt.Println("Entering House..")
}

func closeDoor() {
	fmt.Println("Closing the Door")
}
