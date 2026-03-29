import (
	"slices"
	"maps"
	"cmp"
)


func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	ans := make([]int, 0)

	for _, x := range nums {
		if _, ok := seen[x]; ok {
			seen[x]++
		} else {
			seen[x] = 1
		}
	}

	keys := slices.SortedFunc(maps.Keys(seen), func(a, b int) int {
		return cmp.Compare(seen[b], seen[a])
	})

	for i := range k {
		ans = append(ans, keys[i])
	}

	return ans
}
