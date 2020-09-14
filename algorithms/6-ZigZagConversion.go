将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);
示例 1:
输入: s = "LEETCODEISHIRING", numRows = 3
输出: "LCIRETOESIIGEDHN"

示例 2:
输入: s = "LEETCODEISHIRING", numRows = 4
输出: "LDREOEIIECIHNTSG"
解释:

L     D     R
E   O E   I I
E C   I H   N
T     S     G

//把字符串中每个字符用下标进行表示后写出Z字形变换，可以发现每一行都是有规律的等差数组
//第一行和最后一行各是一个等差数组，其余行可划分为两个等差数组，所有的等差都是2*n-2，只需确定好等差数组的第一个数即可
func convert(s string, numRows int) string {
    if numRows == 1{
        return s
    }
    var res string
    for i := 0 ; i < numRows ; i++ {
        if i == 0 || i == numRows-1 {
            for j := i ; j < len(s) ; j += 2*numRows - 2 {
                res = res + string(s[j])
            }
        }else{
            m := i
            k := 2*numRows-2-i
            for m < len(s) || k < len(s) {
                if m < len(s){
                    res = res + string(s[m])
                }
                if k < len(s){
                    res = res + string(s[k])
                }
                m += 2*numRows - 2
                k += 2*numRows - 2
            }
        }
    }
    return res
}
