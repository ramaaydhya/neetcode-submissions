func hasDuplicate(nums []int) bool {
    temp := make(map[int]struct{})
    for _, v := range nums {
        if _, ok := temp[v]; ok {
            return true
        } 
        temp[v] = struct{}{}
    }
    return false 
}
