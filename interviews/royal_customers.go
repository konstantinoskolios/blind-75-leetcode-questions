// assume you have two log files day1 and day2
// both of the files contains a unique customerId, pageId and timestamp
// find the royalCustomer
// royal customer is the that user that logged into the system 2/2 days and visit at least 2 unique pages each day.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const log1 = `
1,1
1,1
1,4
2,1
1,1
3,1
4,2
5,2
6,3
6,4
7,3
7,4
8,1
8,5
9,2
9,3
100,1
100,1
100,1
100,1
100,1
`

const log2 = `
1,1
1,200
1,3
7,10
8,2
9,1
9,2
11,1
11,2
12,1
12,26
6,100
6,101
100,1
100,1
100,1
100,1
100,1
`

func main() {

	lines1 := strings.Split(strings.TrimSpace(log1), "\n")
	m1 := make(map[int][]string)

	lines2 := strings.Split(strings.TrimSpace(log2), "\n")
	m2 := make(map[int][]string)

	splitData(lines1, m1)
	splitData(lines2, m2)

	eligibleUsers := make(map[int]int)

	eligibleUserCounter(eligibleUsers, m1)
	eligibleUserCounter(eligibleUsers, m2)

	users := make([]int, 0, 10)
	
	for k, v := range eligibleUsers {
		if v > 1 {
			users = append(users, k)
		}
	}

	fmt.Printf("Royal Users: %v\n", users)

}

func splitData(lines []string, m map[int][]string) {
	for _, line := range lines {
		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			continue
		}

		key, err := strconv.Atoi(parts[0])

		if err != nil {
			continue
		}

		value := parts[1]
		m[key] = append(m[key], value)
	}
}

func eligibleUserCounter(eligibleUser map[int]int, m map[int][]string) {
	for user, pages := range m {
	
		if len(arrayToSet(pages)) >= 2 {
			eligibleUser[user]++
		}
	}
}

func arrayToSet(arr []string) map[string]struct{} {
    set := make(map[string]struct{})
	
    for _, v := range arr {
        set[v] = struct{}{}
    }
	
    return set
}
