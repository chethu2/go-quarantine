package main

import (
	"fmt"
	"strings"
)

// splitter function takes two inputs data and separator
// splits the data with specified separator
// Info: this fucntion has named return values
func splitter(data, separator string) (value1, value2 string) {
	splitted_value := strings.Split(data, separator)
	value1, value2 = splitted_value[0], splitted_value[1]
	// return statement without arguments returns named return values are called 'naked' return
	return
}

func main() {
	data := "username_password"
	separator := "_"
	fmt.Printf("Passing the input %s and separator %s to the splitter function\n", data, separator)
	value1, value2 := splitter(data, separator)
	fmt.Printf("Naked return function returns: %s and %s\n", value1, value2)
}
