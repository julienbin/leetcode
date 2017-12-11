const (
    MAX int = 1<<31 - 1
)

func reverse(x int) int {
    sign := 1
    tmp := x
    if x < 0 {
        sign = -1
        tmp = -x
    }
    list := []int{}
    for tmp > 9 {
        list = append(list, tmp % 10)
        tmp = tmp / 10
    }
    list = append(list, tmp)
    ret := 0
    for i := 0; i < len(list); i ++ {
        ret = 10 * ret + list[i]
        if ret > MAX {
            return 0
        }
    }
    return ret * sign
}
