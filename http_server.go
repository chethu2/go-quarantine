package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	keys := req.URL.Query().Get("tags")
	fmt.Printf("key type %T %v\n", keys, keys)
	result := make(map[string]int)
	json.Unmarshal([]byte(keys), &result)
	fmt.Printf("result type %T %v\n", result, result)
	fmt.Fprintf(w, "success\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Listeing ")
	http.ListenAndServe(":8090", nil)
}
