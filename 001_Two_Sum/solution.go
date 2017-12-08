func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    result := make([]int, 2, 2)
    for i, v := range nums {
        if _, ok := m[target - v]; ok {
            result[0] = m[target - v]
            result[1] = i
            break
        }
        m[v] = i
    }
    return result
}
