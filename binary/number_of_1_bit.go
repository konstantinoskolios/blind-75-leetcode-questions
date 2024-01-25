package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(5))
}

func hammingWeight(num uint32) int {
	s := fmt.Sprintf("%b",num)
	count := 0
	for _,v := range s {
		if  v == '1' {
			count++
		}
	}
	return count
}