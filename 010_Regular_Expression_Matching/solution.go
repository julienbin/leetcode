func isMatch(s string, p string) bool {
    lens := len(s)
    lenp := len(p)
    matched := make([][]bool, lens + 1)
    for i := 0; i <= lens; i ++ {
        matched[i] = make([]bool, lenp + 1)
    }
    matched[0][0] = true
    for i := 0; i <= lens; i ++ {
        for j := 1; j <= lenp; j ++ {
            if j > 1 && p[j - 1] == '*' {
                matched[i][j] = matched[i][j - 2] || (i > 0 && (s[i - 1] == p[j - 2] || p[j - 2] == '.') && matched[i - 1][j])
            } else {
                matched[i][j] = i > 0 && matched[i - 1][j - 1] && (s[i - 1] == p[j - 1] || p[j - 1] == '.')
            }
        }
    }
    return matched[lens][lenp]
}
