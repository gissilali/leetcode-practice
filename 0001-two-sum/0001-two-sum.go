func twoSum(nums []int, target int) []int {
	var indices []int
	indexMap := map[int]int{}
	for i, num := range nums {
		fmt.Println(num)
		diff := target - num
		if val, ok := indexMap[diff]; ok {
			indices = []int{val, i}
			break
		} else {
			indexMap[num] = i
		}
	}

	return indices
}
