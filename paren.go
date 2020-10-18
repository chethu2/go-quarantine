package main

import (
	"fmt"
	"strings"
)

var (
	PARENTHESES = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	OPEN_PARENTHESES  = "({["
	CLOSE_PARENTHESES = ")}]"
)

type Element struct {
	value interface{}
	next  *Element
}

type Stack struct {
	top  *Element
	size int
}

func (stack *Stack) Len() int {
	return stack.size
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &Element{value, stack.top}
	stack.size++
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top.value, stack.top.next
		stack.size--
		return
	}
	return nil
}

func main() {
	is_valid := true
	stack := &Stack{}
	var message string

	fmt.Scanf("%s", &message)

	for _, p := range strings.Split(message, "") {
		if strings.Index(OPEN_PARENTHESES, p) != -1 {
			stack.Push(p)
			continue
		}
		if strings.Index(CLOSE_PARENTHESES, p) != -1 {
			if stack.Len() > 0 {
				last_parentheses := stack.Pop().(string)
				if PARENTHESES[last_parentheses] == p {
					continue
				}
			}
			is_valid = false
			break
		}
	}

	if is_valid && stack.Len() == 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
