func romanToInt(s string) int {
    val := map[rune]int{ 
        rune('I') : 1,
        rune('V') : 5,
        rune('X') : 10,
        rune('L') : 50,
        rune('C') : 100,
        rune('D') : 500,
        rune('M') : 1000,
    }
    runes := []rune(s)
    ret := 0
    for i := 0; i < len(runes); i ++ {
        if i > 0 && val[runes[i]] > val[runes[i - 1]] {
            ret += val[runes[i]] - 2 * val[runes[i - 1]]
        } else {
            ret += val[runes[i]]
        }
    }
    return ret
}
