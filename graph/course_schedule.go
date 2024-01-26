package main

type Solution struct{}

func canFinish(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int][]int, numCourses)

	for _, pre := range prerequisites {
		crs, preReq := pre[0], pre[1]
		preMap[crs] = append(preMap[crs], preReq)
	}

	visitSet := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(crs int) bool {
		if visitSet[crs] {
			return false
		}
		if len(preMap[crs]) == 0 {
			return true
		}

		visitSet[crs] = true
		for _, pre := range preMap[crs] {
			if !dfs(pre) {
				return false
			}
		}
		visitSet[crs] = false
		preMap[crs] = nil
		return true
	}

	for crs := range preMap {
		if !dfs(crs) {
			return false
		}
	}

	return true
}

func main() {
	// Example usage
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}


	result := canFinish(numCourses, prerequisites)
	println(result)
}
