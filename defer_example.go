package main

import "fmt"

func main() {
	test()
}

func test() {
	var err error
	fmt.Println("Defer Example : Start")
	defer deferImpl(err)
	fmt.Println("Defer Example: Faking Error")
	err = fakeError("1")
	fmt.Println("Defer Example: End")
}

func deferImpl(err error) {
	if err != nil {
		fmt.Println("Program terminated with error: ", err)
	}
	fmt.Print("Success")
}

func fakeError(id string) error {
	return fmt.Errorf("Error!!: %s", id)
}
