func twoSum(nums []int, target int) []int {
	var indices []int
	currentIndex := 0

	for currentIndex < len(nums) {
		currentNum := nums[currentIndex]
		sum := 0
		for i, num := range nums {
			if currentIndex != i {
				sum = currentNum + num
			} else {
				continue
			}

			if sum == target {
				indices = []int{currentIndex, i}
				break
			}
		}

		if sum == target {
			break
		}

		currentIndex++
	}

	return indices
}
