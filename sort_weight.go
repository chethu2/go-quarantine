package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var outputSlice []int
	//var stringInput string
	// reading the input from STDIN
	//fmt.Scanf("%v", &stringInput)

	for _, v := range strings.Fields("99 100") {
		fmt.Println("v", v)
		intValue := getSum(v)
		outputSlice = append(outputSlice, intValue)
	}
	sort.Ints(outputSlice)
	stringOutput := getOutputString(outputSlice)
	fmt.Print(stringOutput)
}

func getSum(stringValue string) int {
	sum := 0
	i, _ := strconv.Atoi(stringValue)
	for i > 10 {
		sum = sum + i%10
		i = i / 10
		sum = sum + i
	}
	return sum + i
}

func getOutputString(intSlice []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intSlice)), " "), "[]")
}
