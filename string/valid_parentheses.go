package main

import "fmt"

func main() {
	s1 := "{}{}(){}[]{[()]}"
	fmt.Printf("%v\n", isValid(s1))
}

func isValid(s string) bool {
    size := len(s)

    if size%2 != 0 {
        return false
    }

    stack := []rune{}

    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 {
                return false
            }

            lastOpen := stack[len(stack)-1]
			fmt.Printf("last open: %c\n",lastOpen)
            stack = stack[:len(stack)-1]
			fmt.Printf("stack: %c\n",stack)

            if char == ')' && lastOpen != '(' ||
               char == '}' && lastOpen != '{' ||
               char == ']' && lastOpen != '[' {
                return false
            }
        }
    }

    return len(stack) == 0
}
