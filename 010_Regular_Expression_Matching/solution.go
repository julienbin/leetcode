func isMatch(s string, p string) bool {
    lens := len(s)
    lenp := len(p)
    
    // initial matched matrix
    matched := make([][]bool, lenp + 1)
    for i := 0; i <= lenp; i ++ {
        matched[i] = make([]bool, lens + 1)
    }
    matched[0][0] = true
    
    // first column
    for i := 2; i <= lenp; i += 2 {
        if p[i - 1] == '*' {
            matched[i][0] = true
        } else {
            break
        }
    }
    
    // compute the matrix
    for i := 1; i <= lenp; i ++ {
        if p[i - 1] == '*' {
            for j := 1; j <= lens; j++ {
                matched[i][j] = matched[i - 2][j] || ((p[i - 2] == '.' || p[i - 2] == s[j - 1]) && matched[i][j - 1])
            }        
        } else {
            for j := 1; j <= lens; j++ {
                matched[i][j] = matched[i - 1][j - 1] && (p[i - 1] == '.' || p[i - 1] == s[j - 1])
            }
        }
    }
    return matched[lenp][lens]
}
