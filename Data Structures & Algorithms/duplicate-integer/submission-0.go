func hasDuplicate(nums []int) bool {
    numMap := map[int]struct{}{}
    for _, num := range nums{
        if _, ok := numMap[num];ok{
            return true
        }
        numMap[num] = struct{}{}
    }
    return false
}
