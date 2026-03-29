func isAnagram(s string, t string) bool {
	symbols := make(map[rune]int)

	for _, x := range s {
		_, ok := symbols[x]
		if ok {
			symbols[x]++
		} else {
			symbols[x] = 1
		}
	}  

	for _, y := range t {
		_, ok := symbols[y]
		if ok {
			symbols[y]--
		} else {
			return false
		} 
	}

	for _, v := range symbols {
		if v != 0 {
			return false
		}
	}

	return true
}
