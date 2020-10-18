package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	outPutMap := make(map[int]int)
	var outputSlice []int
	in := bufio.NewReader(os.Stdin)

	stringInput, _ := in.ReadString('\n')
	testArray := strings.Fields(stringInput)
	for index, v := range testArray {
		intValue := getSum(v)
		outPutMap[intValue] = index
		outputSlice = append(outputSlice, intValue)
	}
	sort.Ints(outputSlice)
	stringOutput := getOutputString(outputSlice, testArray, outPutMap)
	fmt.Print(stringOutput)
}

func getOutputString(intSlice []int, testArray []string, outputMap map[int]string) string {
	weightString := ""
	for _, v := range intSlice {
		weightString = weightString + " " + testArray[outputMap[v]]
	}
	return weightString
}

func getSum(stringValue string) int {
	sum := 0
	i, _ := strconv.Atoi(stringValue)
	for i > 10 {
		sum = sum + i%10
		i = i / 10
	}
	return sum + i
}
