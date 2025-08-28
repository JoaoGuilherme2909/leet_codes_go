package twosum

func TwoSum(nums []int, target int) []int {
	checkedNums := make(map[int]int)

	for i, n := range nums {
		val, ok := checkedNums[n]
		if ok {
			return []int{val, i}
		}

		num := target - n
		checkedNums[num] = i
	}

	return nil
}
