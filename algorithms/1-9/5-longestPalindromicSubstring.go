给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"

//分别以s[i]为中心向两边进行扩展找回文子串，奇偶情况都要考虑到，遍历结束后，找到的最长子串就是S中的最长回文子串

func longestPalindrome(s string) string {
    var res string
    for i := 0 ; i < len(s) ; i++ {
        //奇数情况
        l := i - 1 
        r := i + 1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            l --
            r ++
        }
        if len(res) < r-l-1 {
            res = s[l+1:r]
        }
        //偶数情况
        l = i
        r = i+1
        for l >= 0 && r < len(s) && s[l] == s[r] {
            l --
            r ++
        }
        if len(res) < r-l-1 {
            res = s[l+1:r]
        }

    }
    return res
}
