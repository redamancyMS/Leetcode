给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


//可以采用滑动窗口的思想，当window[c] 大于1的时候，说明窗口中存在重复的字符，不符合条件，就应该移动left来缩小窗口

func lengthOfLongestSubstring(s string) int {
    window := make(map[byte]int) 

    right := 0
    left := 0
    length := 0

    for right < len(s) {
        c := s[right]
        right ++
        window[c]++
        for window[c] > 1 {
            d := s[left]
            left ++
            window[d]--
        }
        length = max(length,right-left)
    }
    return length
}

func max(a,b int) int {
    if a > b {
        return a 
    }
    return b
}
