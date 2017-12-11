const (
    MAX int = 1<<31 - 1
    MIN int = -1 * MAX - 1
)

func myAtoi(str string) int {
    if len(str) == 0 {
        return 0
    }
    runes := []rune(str)
    sign := 1
    i := 0
    for ; runes[i] == rune(' '); i ++ {
    }
    
    if runes[i] == rune('-') {
        sign = -1
        i ++
    } else if runes[i] == rune('+') {
        i ++
    }
    
    ret := 0
    for ; i < len(str); i ++ {
        val := int(runes[i] - rune('0'))
        if val < 0 || val > 9 {
            break;
        } else {
            ret = 10 * ret + val
        }
        if ret * sign > MAX {
            return MAX
        }
        if ret * sign < MIN {
            return MIN
        }
    }
    return ret * sign
}
