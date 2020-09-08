package lengthOfLastWord

func lengthOfLastWord(s string) int {
    if len(s) == 0 {
        return 0
    }
    spaceFlag := 0
    length := 0
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' '{
            if spaceFlag > 0 {
                 break
            }
        }else {
            length += 1
            spaceFlag += 1
        }
    }
    return length
}
