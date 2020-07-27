func threeSum(nums []int) [][]int {
    // 首先排序
    // [-1,0,1,2,-1,-4]
    // [-4, -1, -1, 0, 1, 2]
    result := [][]int{} 
    sort.Ints(nums)
    // fmt.Println("nums: ", nums)
    for i := 0; i < len(nums); i++{
        // 如果第一个数大于0，那么和不可能等于0，直接退出
        if nums[i] > 0 {
            break
        }
        // 如果当前的nums[i]与前一个数相等，则退出当前循环
        if i > 0 && nums[i] == nums[i - 1]{
            continue
        }
        left := i + 1
        right := len(nums) - 1
        for left < right{
            // 左边下标大于右边下标直接退出
            // if left >= right{
            //     break
            // }
            // 如果左下标与上一位置相等，则左下标右移一位
            if left > i + 1 && nums[left] ==nums[left - 1]{
                left++
                continue
            }
            // 如果右下标与上一位置相等，则左下标右移一位
            if right < len(nums) - 1 && nums[right] ==nums[right + 1]{
                right--
                continue
            }
            // fmt.Printf("i=%d  left=%d  right=%d\n", i, left, right)
            sum := nums[left] + nums[right] + nums[i]
            // 如果三数和等于0,保存结果；同时左下标右移一位，右下标左移一位
            if sum == 0{
                threeSumList := []int{nums[i], nums[left], nums[right]}
                result = append(result, threeSumList)
                left++
                right--
            // 如果sum大于0，说明sum需要减少，右下标左移
            }else if sum > 0 {
                right--
            // 如果sum小于0，说明sum需要增加，左下标右移
            }else{
                left++
            }
        }
    }
    return result
}
