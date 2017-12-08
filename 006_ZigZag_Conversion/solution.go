func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    runes := []rune(s)
    rows := make([][]rune, numRows)
    i := 0
    step := 1
    for index := range runes {
        rows[i] = append(rows[i], runes[index])
        if i == numRows - 1 {
            step = -1
        } else if i == 0 {
            step = 1
        }
        i += step
    }
    retRunes := []rune{}
    for index := range rows {
        retRunes = append(retRunes, rows[index]...)
    }
    return string(retRunes)
}
