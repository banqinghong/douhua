package twoSum

func twoSum(nums []int, target int) []int {
    // 新建差值map，用于存放差值
    diffMap := make(map[int]int)
    for i := 0; i < len(nums); i++{
        // 判断差值是否已经在map中存在,存在则返回，不存在继续循环计算
        index, ok := diffMap[target - nums[i]]
        if ok{
            return []int{index, i}
        }
        diffMap[nums[i]] = i
    }
    return nil
}
