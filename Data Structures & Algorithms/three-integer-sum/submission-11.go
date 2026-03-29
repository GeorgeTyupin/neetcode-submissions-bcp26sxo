import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	ans := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
		}
	}

	return removeDuplicates(ans)
}

func removeDuplicates(nums [][]int) [][]int {
	lookup := make(map[string]struct{})
	result := make([][]int, 0)

	for _, v := range nums {
		key := fmt.Sprint(v)
		if _, ok := lookup[key]; !ok {
			lookup[key] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}