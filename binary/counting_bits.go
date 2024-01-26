package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(countBits(2))
}


func countBits(n int) []int {

	arr := make([]int, n+1)

	for i:=0; i <=n;i++ {
		arr[i] = bits.OnesCount(uint(i))
	}
	
	return arr
}
