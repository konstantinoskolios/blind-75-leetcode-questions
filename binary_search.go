package main

import "fmt"

var x = []int{1,2,5,12}

func main() {
	findMe(1, 0, len(x)-1)
}

func findMe(t, s, e int) {
	if s > e {
		fmt.Println("not found")
		return
	}

	m := (s + e) / 2

	if x[m] == t {
		fmt.Println("found")
		return
	}

	if x[m] > t {
		findMe(t, s, m-1)
	} else {
		fmt.Printf("%v %v %v \n",t,x[m+1],x[e])
		findMe(t, m+1, e)
	}
}
