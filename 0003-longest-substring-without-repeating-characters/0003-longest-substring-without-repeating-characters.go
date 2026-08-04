func lengthOfLongestSubstring(s string) int {
    letterMap := make(map[rune]struct{})
    longest := 0
    current := 0
    start := 0

    runes := []rune(s)

    for i := 0; i < len(runes); i++ {
        value := runes[i]

        for {
            _, ok := letterMap[value]
            if !ok {
                break
            }
            delete(letterMap, runes[start])
            start++
        }

        letterMap[value] = struct{}{}
        current = i - start + 1

        if current >= longest {
            longest = current
        }
    }

    return longest
}