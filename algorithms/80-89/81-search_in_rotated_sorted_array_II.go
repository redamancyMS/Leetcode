//这个题目不管怎样最坏的时间复杂度都是O(n)的
func search(nums []int, target int) bool {
    for i := 0 ; i < len(nums) ; i ++ {
        if nums[i] == target {
            return true
        }
    }
    return false
}

//将数组末尾与数组头元素相同的元素删除
func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    l,r := 0,len(nums)-1
    for r > l && nums[r] == nums[l] {
        r--
    }
    k := r
    for l < r {
        mid := l + r + 1 >> 1
        if nums[mid] >= nums[l] {
            l = mid
        }else{
            r = mid-1
        }
    }
    if target >= nums[0]{
        l = 0
    }else{
        l = r + 1
        r = k
    }
    for l < r {
        mid := (l + r )/2
        if nums[mid] >= target {
            r = mid
        }else{
            l = mid+1
        }
    }
    if nums[r] == target {
        return true
    }
    return false
}