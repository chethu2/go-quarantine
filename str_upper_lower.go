package main

import (
	"fmt"
	"unicode"
)

// 1. read input i.e, string
// 2. iterate over the string input
// 3. I am using the `unicode` library
// 4. check if character is upper, if yes convert to lower
// 5. check if character is lower, if yes convert to upper

func main() {
	// declare the variable
	var input string

	// read the value
	fmt.Scanf("%s", &input)

	// convert input string into rune slice
	output := []rune(input)

	// iterate over the string input and
	// check for charater case and change accordingly.
	for i := 0; i < len(output); i++ {
		if unicode.IsUpper(output[i]) {
			output[i] = unicode.ToLower(output[i])
		} else if unicode.IsLower(output[i]) {
			output[i] = unicode.ToUpper(output[i])
		}
	}
	// output convert rune slice into string
	fmt.Println(string(output))
}
