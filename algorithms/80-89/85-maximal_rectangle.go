//f(i,j)表示（i,j）这个格子往上数有多少个1 
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    n := len(matrix)
    m := len(matrix[0])
    var h [][]int = make([][]int,n)
    for i := 0 ; i < n ; i++ {
        h[i] = make([]int,m)
    }
    for i := 0 ; i < n ; i++ {
        for j := 0 ; j < m ; j ++ {
            if matrix[i][j] == '1'{
                if i > 0 {
                    h[i][j] = h[i-1][j] + 1
                }else{
                    h[i][j] = 1
                }
            }
        }
    }
    var largestRectangleArea func([]int) int
largestRectangleArea = func(heights []int) int{
    n := len(heights)
    if n == 0 {
        return 0
    }
    var stack1 []int
    var stack []int
    var left []int = make([]int,n)
    var right []int = make([]int,n)
    ans := 0 

    for i := 0 ; i < n ; i ++ {
        for len(stack1) > 0 && heights[stack1[len(stack1)-1]] >= heights[i]{
            stack1 = stack1[:len(stack1)-1]
        }
        if len(stack1) == 0 {
            left[i] = -1
        }else{
            left[i] = stack1[len(stack1)-1]
        }
        stack1 = append(stack1,i)
    }

    for i := n-1 ; i >= 0 ; i-- {
        for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
            stack = stack[:len(stack)-1]
        }
        if len(stack) == 0 {
            right[i] = n 
        }else{
            right[i] = stack[len(stack)-1]
        }
        stack = append(stack, i)
    }
    for i := 0 ; i < n ; i++ {
        ans = max(ans,heights[i]*(right[i]-left[i]-1))
    }

    return ans  
    
}


    res := 0
    for i := 0 ; i < n ; i++ {
        res = max(res, largestRectangleArea(h[i]))
    }
    return res
}

func max(a,b int) int {
    if a >b {
        return a
    }
    return b
}