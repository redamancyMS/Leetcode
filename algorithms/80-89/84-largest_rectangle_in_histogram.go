//单调栈的经典用法，找每个数左边第一个比它小的数以及右边第一个比它小的数
//这里面栈存储的是数组的下标
func largestRectangleArea(heights []int) int {
    n := len(heights)
    var stack1,stack2 []int
    var left []int = make([]int,n)
    var right []int = make([]int,n)
    res := 0
    for i := 0 ; i < n ; i++{
        for len(stack1) > 0 && heights[stack1[len(stack1)-1]] >= heights[i] {
            stack1 = stack1[:len(stack1)-1]
        }  
        if len(stack1) == 0 {
            left[i] = -1
        }else{
            left[i] = stack1[len(stack1)-1]
        }
        stack1 = append(stack1,i)
    }
    for i := n-1 ; i >=0 ; i-- {
        for len(stack2) > 0 && heights[stack2[len(stack2)-1]] >= heights[i] {
            stack2 = stack2[:len(stack2)-1]
        }
        if len(stack2) == 0 {
            right[i] = n
        }else{
            right[i] = stack2[len(stack2)-1]
        }
        stack2 = append(stack2,i)
    }
    for i := 0 ; i < n ; i ++{
        res = max(res,heights[i]*(right[i]-left[i]-1))
    }
    return res

}
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}