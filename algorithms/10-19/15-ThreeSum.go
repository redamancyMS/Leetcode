给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

//固定三个元素中的一个，剩下两个元素采用双指针的方法进行遍历
func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)
    n := len(nums)
    for i := 0 ; i < n ; i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        target := -1 * nums[i]
        k := n - 1
        for j := i + 1 ; j < n ; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            for j < k && nums[j]+nums[k] > target {
                k --
            }
            if j==k {
                break
            }
            if nums[j]+nums[k]==target{
                //res[m] = []int{nums[i],nums[j],nums[k]}
                res = append(res,[]int{nums[i],nums[j],nums[k]})
            }
        }
    }
    return res
}
