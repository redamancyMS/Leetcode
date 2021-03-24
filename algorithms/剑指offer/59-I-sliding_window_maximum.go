func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    var ans []int
    if n < 1 {
        return ans
    }
    var queue []int = make([]int,n)
    head := 0
    tail := -1
    for i := 0 ; i < n ; i++ {
        if head <= tail && i-k+1 > queue[head]{
            head++
        }
        for head <= tail && nums[queue[tail]] < nums[i]{
            tail--
        }
        tail++
        queue[tail] = i
        if i-k+1 >= 0 {
            ans = append(ans,nums[queue[head]])
        }

    }
    return ans
}