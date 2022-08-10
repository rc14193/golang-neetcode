func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, value := range nums{
        _, ok := m[value]
        if(ok){
            return true
        }else{
            m[value] = true
        }
    }
    return false
}
