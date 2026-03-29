import "slices"

func checkInclusion(s1 string, s2 string) bool {
	c1, c2 := []byte(s1), []byte(s2)
	slices.Sort(c1)
	l, r := 0, len(c1)
	win := c2[l:r]

	for r <= len(c2) {
		if compareStrings(c1, win) {
			return true
		} else {
			l++
			r++
			win = c2[l:r]
		}
	}

	return false
}

func compareStrings(c1, win []byte) bool {
	subWin := slices.Clone(win)
	slices.Sort(subWin)
	return slices.Equal(subWin, c1)
}