func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    i, j := 0, 0
    nums3 := []int{}
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] < nums2[j] {
            nums3 = append(nums3, nums1[i])
            i = i + 1
        } else {
            nums3 = append(nums3, nums2[j])
            j = j + 1
        }
    }
    if i < len(nums1) {
        for ; i < len(nums1); i++ {
            nums3 = append(nums3, nums1[i])
        }
    }
    if j < len(nums2) {
        for ; j < len(nums2); j++ {
            nums3 = append(nums3, nums2[j])
        }
    }
    lenSum := len(nums3)
    if lenSum % 2 == 0 {
        return (float64(nums3[lenSum / 2 - 1]) + float64(nums3[lenSum / 2])) / 2.0
    } else {
        return float64(nums3[(lenSum + 1) / 2 - 1])
    }
}
