编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

//用双重循环判断数组中的每个字符串对应的前缀字母是否相同
//双重循环解决这个问题
func longestCommonPrefix(strs []string) string {
    var res string
    if len(strs) == 0 {
        return res
    }

    for i := 0 ;; i++ {
        if i >= len(strs[0]) {
            return res
        }
        c := strs[0][i]
        for _,str := range strs {
            if i >= len(str) || str[i] != c{
                return res
            }
        }
        res += string(c)
    } 
    return res
}
