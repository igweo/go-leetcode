package main

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool, len(nums))
	for _, n := range nums {
		if visited[n] == true {
			return visited[n]
		} else {
			visited[n] = true
		}
	}
	return false
}

// O(N) time complexity
// O(N) space complexity
