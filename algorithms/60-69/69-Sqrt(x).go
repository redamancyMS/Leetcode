//这是一个二分的问题，找到一个最大的y，让y的平方小于等于x
func mySqrt(x int) int {
    l := 0
    r := x
   ans := -1
    for l <= r {
        mid := l + (r-l)/2
        if (mid*mid <= x){
            ans = mid
            l = mid+1
        }else {
            r = mid-1
        }
    }
    return ans
}