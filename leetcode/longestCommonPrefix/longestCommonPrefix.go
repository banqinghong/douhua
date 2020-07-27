func longestCommonPrefix(strs []string) string {
    // 长度为0的时候直接返回空字符串
	if len(strs) == 0 {
        return ""
    }
    // 以列表第一个字符串为基准字符串， 逐一判断每一个字符
    for index, char := range strs[0] {
        // 根据基准字符串对比列表中其他字符串相同位置字符是否相等
        for _, v := range strs {
            // 如果index大于当前字符串长度，直接返回当前字符串。不可能还有更长的公共前缀
            if index >= len(v) {
                return v
            }
            // 如果当前字符串不相等，说明前面的是已经匹配上的
            if v[index] != byte(char) {
                return v[:index]
            }
        }
    }
    return strs[0]
}

