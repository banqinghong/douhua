package reverseInteger

import "math"

func reverse(x int) int {
    // 2^31 = 2147483648
    limit_max := int((math.Pow(2,31) - 1) / 10) // 
    limit_min := int(- math.Pow(2,31) / 10)
    ret := 0
    for x != 0 {
        pop := x % 10
        if ret > limit_max || (ret == limit_max && pop > 7){
            return 0
        }
        if ret < limit_min || (ret == limit_min && pop < -8){
            return 0
        }
        ret = ret * 10 + pop
        x = x / 10
    }
    return ret
}
