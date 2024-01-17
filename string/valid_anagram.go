package main

import (
	"fmt"
	"sort"
)


func main(){
	fmt.Printf("%v\n",isAnagram("aca","caa"))
}



func isAnagram(s string, t string) bool {
    return sortedString(s) == sortedString(t)
}

func sortedString(s string) string {
    runes := []rune(s)
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })
    return string(runes)
}