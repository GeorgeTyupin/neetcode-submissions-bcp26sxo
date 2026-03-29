func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

	sMap := make(map[rune]int)
	tMap := make(map[rune]int)

	for _, x := range []rune(s) {
		_, ok := sMap[x]
		if ok {
			sMap[x]++
		} else {
			sMap[x] = 1
		}
	}

	for _, x := range []rune(t) {
		_, ok := tMap[x]
		if ok {
			tMap[x]++
		} else {
			tMap[x] = 1
		}
	}

    for k, v1 := range sMap {
        v2, ok := tMap[k]
        if !ok || v1 != v2 {
            return false
        }
    }

    return true
}
