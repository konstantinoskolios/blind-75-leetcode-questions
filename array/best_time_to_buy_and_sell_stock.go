package main

import (
	"fmt"
)

var prices1 = []int{7, 1, 5, 3, 6, 4}
var prices2 = []int{7, 6, 4, 3, 1}
var prices3 = []int{2, 4, 1}

func main() {
	fmt.Printf("result: %v \n", maxProfit(prices1))
	fmt.Printf("result: %v \n", maxProfit(prices2))
	fmt.Printf("result: %v \n", maxProfit(prices3))
}
func maxProfit(prices []int) int {

    var minPrice = prices[0]
    var profit = 0

    if len(prices) > 1 {

        for pos := range prices {

            if prices[pos] < minPrice {
                minPrice = prices[pos]
            }

            if prices[pos]-minPrice > profit {
                profit = prices[pos] - minPrice
            }
        }
    }

    return profit
}
