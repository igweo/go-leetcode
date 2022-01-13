package main

func missingNumber(nums []int) int {
	incompleteArr := 0
	completeArr := 0
	for _, n := range nums {
		incompleteArr += n
	}

	for i := 0; i < len(nums)+1; i++ {
		completeArr += i
	}

	return completeArr - incompleteArr
}

// O(N) time complexity
// O(1) space complexity
