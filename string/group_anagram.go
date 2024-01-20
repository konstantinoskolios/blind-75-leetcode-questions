package main

import (
	"fmt"
	"sort"
)

func main() {

	groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"})

}

func groupAnagrams(strs []string) [][]string {
    
	hm := make(map[string][]string)
	
	for i:=0;i<len(strs);i++{
		ov := strs[i]
		sv := sortedString(strs[i])
		if _,ok := hm[sv]; !ok {
			hm[sv] = []string{ov}
		}else{
			hm[sv] = append(hm[sv],ov)
		}
	}
	
	arr := [][]string{}

	for k := range hm {
		arr = append(arr, hm[k])
	}

	return arr

	
}

func sortedString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}