package main
func main()  {
	
}
func strStr(haystack string, needle string) int {

    // Lengths of haystack and needle
    n := len(haystack)
    m := len(needle)

    // If needle is longer than haystack, it cannot exist
    if m > n {
        return -1
    }

    // Try matching needle starting from each valid index in haystack
    for i := 0; i <= n-m; i++ {

        j := 0

        // Compare characters of needle with haystack
        for j < m && haystack[i+j] == needle[j] {
            j++
        }

        // If we matched all characters of needle
        if j == m {
            return i
        }
    }

    // If needle is not found
    return -1
}
