func searchRange(nums []int, target int) []int {
    var res  = []int{-1,-1}
    if len(nums) == 0 {
        return res
    }
    
    l := 0
    r := len(nums) - 1
    for l < r {
        mid := l + (r-l) / 2
        if nums[mid] >= target {
            r = mid
        }else {
            l = mid + 1
        }
    }
    if nums[r] != target {
        return res
    }
    res[0] = r
    
   l = 0
   r = len(nums) - 1
    for l < r {
        mid := l + r + 1 >> 1
        if nums[mid] <= target {
            l = mid 
        }else {
            r = mid - 1 
        }
    }
    res[1] = r
    return res
}