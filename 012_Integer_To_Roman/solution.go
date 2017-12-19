func intToRoman(num int) string {
    symbol := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    value := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    str := ""
    temp := num
    for i := 0; temp != 0; i ++ {
        for temp >= value[i] {
            temp -= value[i]
            str += symbol[i]
        }
    }
    return str
}
