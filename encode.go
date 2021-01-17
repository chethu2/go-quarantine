package main

import (
	"fmt"
	"net/url"
)

func main() {

	// Use net/url's value list to build query, then encode it
	params := url.Values{}
	params.Add("q", "1 + 2")
	params.Add("s", "example for golangcode.com")
	output := params.Encode()

	fmt.Println("After:", output)
}
