//和三数之和的原理相同，将循环降低一重
func fourSum(nums []int, target int) [][]int {
    var res [][]int
    if len(nums) < 4 {
        return res
    }

    sort.Ints(nums)
    for i := 0 ; i < len(nums) ; i ++ {
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        for j := i + 1 ; j < len(nums) ; j++ {
            if j > i + 1 && nums[j] == nums[j-1]{
                continue
            }
            r := len(nums)-1 
            for l := j+1 ; l < r ; l++ {
                if l > j+1 && nums[l]==nums[l-1]{
                    continue
                }
                for r-1 > l && nums[i]+nums[j]+nums[l]+nums[r-1] >=target{
                    r--
                }
                if nums[i]+nums[j]+nums[l]+nums[r] ==target {
                    res = append(res,[]int{nums[i],nums[j],nums[l],nums[r]})
                }
            }
        }
    }
    return res
}
