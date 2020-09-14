判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

//这个题目可以参考数字反转的做法，将x重新提取，进行整数反转，判断反转后的整数和原来的整数是否相同
//如果想取巧的话也可以将数字直接转换成字符串进行判断

func isPalindrome(x int) bool {
     if x < 0 {
         return false
     }

     if x < 10 {
         return true
     }

     res := 0
     temp := x
     for temp != 0 {
         res = res * 10 + temp % 10
         temp = temp / 10
     }

     if res == x {
         return true
     }else{
         return false
     }
     
}
