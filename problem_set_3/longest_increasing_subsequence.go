package main

import "fmt"

func lengthOfLIS(nums []int) int {
	end := len(nums)
	database := make([]int, end)

    for i := range database {
        database[i] = 1
    }

	for idx := 1; idx < end; idx++ {
		for counter := 0; counter < idx; counter++ {
			if nums[idx] > nums[counter] {
				database[idx] = max(database[idx], database[counter]+1)
			}
		}
	}

	return max(database...)
}

func max(nums ...int) int {
	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{1, 1, 1, 1, 1, 1, 1}))
}
