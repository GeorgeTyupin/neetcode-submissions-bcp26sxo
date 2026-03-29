import "slices"

func groupAnagrams(strs []string) [][]string {
	out := make([][]string, 0)
	lookup := make(map[string][]string)

	for _, s := range strs {
		key := getKey(s)
		lookup[key] = append(lookup[key], s)
	}

	for _, v := range lookup {
		out = append(out, v)
	}

	return out
}


func getKey(s string) string {
	b := []byte(s)
	slices.Sort(b)
	return string(b)
}