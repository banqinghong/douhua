func isPalindrome(x int) bool {
    if x == 0 {
        return true
    }
    if x < 0 {
        return false
    }
    if x % 10 == 0 {
        return false
    }
    res := 0
    for res <= x {
        quotient := x / 10
        res = res * 10 + x % 10
        if res == x || res == quotient {
            return true
        }
        x = quotient
    }
    return false
}
