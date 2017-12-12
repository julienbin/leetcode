const (
    MAX int = 1<<31 - 1
)
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    tmp := x
    reverse := 0
    for ; tmp > 0; {
        reverse = 10 * reverse + tmp % 10
        if reverse > MAX {
            return false
        }
        tmp = tmp / 10
    }
    return (reverse == x)
}
