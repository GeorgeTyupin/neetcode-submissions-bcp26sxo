

func hasDuplicate(nums []int) bool {
    dups := make(map[int]struct{})

    for _, x := range nums{
        _, ok := dups[x]
        dups[x] = struct{}{}

        if ok {
            return true
        }
    }

    return false
}