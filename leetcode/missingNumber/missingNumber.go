package missingNumber

func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1
	for {
		// fmt.Println("left= ", left, " right= ", right)
		if left >= right {
			break
		}
		mid := ( left + right ) / 2
		if nums[mid] == mid {
			left = mid + 1
		}else {
			right = mid
		}
	}
	if right != nums[right]{
		return right
	}else {
		return len(nums)
	}
}


