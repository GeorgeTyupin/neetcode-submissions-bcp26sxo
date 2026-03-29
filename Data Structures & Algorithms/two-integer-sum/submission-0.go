func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, x := range nums {
        idx, ok := m[x] 
        if ok{
            return []int{idx, i}
        }
		m[target-x] = i
	}

	return []int{}
}
