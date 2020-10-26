//可以用二进制的方式来记录每一位的数是不是被选择
func subsets(nums []int) [][]int {
    var res [][]int
    n := len(nums)
    for i := 0 ; i < (1 << n ) ; i++ {
        var path []int
        for j := 0 ; j < n ; j++ {
            if (i >> j) & 1 > 0 {
                path = append(path,nums[j])
            }
        }
        res = append(res,path)
    }
    return res
}