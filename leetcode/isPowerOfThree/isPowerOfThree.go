package isPowerOfThree

func isPowerOfThree(n int) bool {
    if n == 1 || n == 3 {
        return true
    }
    // n % 3 余数
    // n / 3 商
    if  n % 3 != 0 || n == 0 {
        return false
    }
    return isPowerOfThree(n / 3)
}
