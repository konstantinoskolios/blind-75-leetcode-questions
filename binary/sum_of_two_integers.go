package main

import "fmt"

func main() {
	fmt.Println(getSum(1,10))
}

func getSum(a int, b int) int {
	for b != 0 {
        carry := a & b

        // Sum of bits of a and b where at least one of the bits is not set
        a = a ^ b

        // Carry is shifted by one so that adding it to a gives the required sum
        b = carry << 1
    }

    return a
}