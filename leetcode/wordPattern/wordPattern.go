package wordPattern

func wordPattern(pattern string, s string) bool {
    // 将pattern 和 s 转化为切片   pattern可以不转换，个人强迫症一定要转换
	patternList := strings.Split(pattern, "")
	sList := strings.Split(s, " ")

    if len(patternList) != len(sList) {
        return false
    }

    // 定义两个map，用来确认双向一致
	sMap := make(map[string]string)
    pMap := make(map[string]string)
	for i := 0; i < len(sList); i++ {
		v, ok := pMap[patternList[i]]
		if ok {
			if v != sList[i] {
				return false
			}
		} else {
			pMap[patternList[i]] = sList[i]
		}

        j, ok := sMap[sList[i]]
		if ok {
			if j != patternList[i] {
				return false
			}
		} else {
			sMap[sList[i]] = patternList[i]
		}
	}
	return true
}