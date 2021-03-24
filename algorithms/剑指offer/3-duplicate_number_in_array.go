func onesCount(x int) (c int) {
    for ; x > 0; x /= 2 {
        c += x % 2
    }
    return
}

func sortByBits(a []int) []int {
    sort.Slice(a, func(i, j int) bool {
        x, y := a[i], a[j]
        cx, cy := onesCount(x), onesCount(y)
        return cx < cy || cx == cy && x < y
    })
    return a
}


// 空间复杂度较上法优化
func findRepeatNumber(nums []int) int {
    n := len(nums)
    for _,num := range nums {
        if num < 0 || num >= n {return -1}
    }
    for i := 0 ; i < n ; i++{
        for i != nums[i] && nums[nums[i]] != nums[i]{
            swap(&nums[i],&nums[nums[i]])
        }
        if i != nums[i] && nums[nums[i]] == nums[i]{return nums[i]}
    }
    return -1
}
func swap(a,b *int) {
    *a,*b = *b,*a
}