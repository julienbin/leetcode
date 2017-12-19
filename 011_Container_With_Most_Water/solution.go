func maxArea(height []int) int {
    area := 0
    left, right := 0, len(height) - 1
    
    for left < right {
        area = max(area, min(height[left], height[right]) * (right - left))
        if height[left] > height[right] {
            right --
        } else {
            left ++
        }
    }
    return area
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}
