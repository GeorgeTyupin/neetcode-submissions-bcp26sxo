import (
	"slices"
	"maps"
	"cmp"
)

func characterReplacement(s string, k int) int {
	l, r := 0, -1
	result := 0
	popularChar := make(map[byte]int)

	for l < len(s) {
		for r+1 < len(s) {
			popularChar[s[r+1]]++
			if isValidWin(r+1-l+1, k, popularChar) {
				r++
			} else {
				popularChar[s[r+1]]--
				break
			}
		}
		if r-l+1 > result {
			result = r - l + 1
		}
		popularChar[s[l]]--
		l++
	}

	return result
}

func isValidWin(lenWin, k int, m map[byte]int) bool {
	return lenWin-getMaxCountMapElem(m) <= k
}

func getMaxCountMapElem(m map[byte]int) int {
	keys := slices.SortedFunc(maps.Keys(m), func(a, b byte) int {
		return cmp.Compare(m[b], m[a])
	})

	return m[keys[0]]
}

