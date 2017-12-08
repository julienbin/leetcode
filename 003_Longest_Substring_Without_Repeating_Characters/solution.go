func lengthOfLongestSubstring(s string) int {
    byteIndexMap := map[byte]int{}
    max := 0
    for i := range s {
        _, ok := byteIndexMap[s[i]]
        if !ok {
            byteIndexMap[s[i]] = i
            if len(byteIndexMap) > max {
                max = len(byteIndexMap)
            }
        } else {
            old := byteIndexMap[s[i]]
            for key, value := range byteIndexMap {
                if value < old {
                    delete(byteIndexMap, key)
                }
            }
            byteIndexMap[s[i]] = i  
        }
    }
    return max
}
