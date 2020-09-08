package containsDuplicate

func containsDuplicate(nums []int) bool {
    numMap := make(map[int]int)
    for k, v := range nums {
        if _, ok := numMap[v]; ok{
            return true
        }else {
            numMap[v] = k
        }
    }
    return false
}
