package main

import (
	"fmt"
)

func main() {
	var num uint32 = 0b11111111111111111111111111111101
	reversed := reverseBits(num)
	fmt.Printf("Original: %032b\nReversed: %032b\n", num, reversed)


}

func reverseBits(num uint32) uint32 {
    var reversed uint32
    for i := 0; i < 32; i++ {
        reversed = (reversed << 1) | (num & 1)
        num >>= 1
    }
    return reversed
}
