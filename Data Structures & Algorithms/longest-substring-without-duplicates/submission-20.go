import "slices"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	dups := make([]byte, 0)
	l, res := 0, 1

	for r := 0; r < len(s); {
		if !slices.Contains(dups, s[r]) {
			dups = append(dups, s[r])
			r++
			continue
		}

		if res < len(dups) {
			res = len(dups)
		}

		for slices.Contains(dups, s[r]) {
			dups = dups[1:]
			l++
		}
	}

	if res < len(dups) {
		res = len(dups)
	}

	return res
}





