func containsDuplicate(nums []int) bool {
	var duplicateFreeArray []int
	for _, num := range nums {
		if contains(duplicateFreeArray, num) {
			return true
		}
		
		duplicateFreeArray = append(duplicateFreeArray, num)
	}

	return false
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}