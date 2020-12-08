//由1到n形成的二叉搜索树的个数是固定的，是卡特兰数C(下:2n,上:n)/(n+1)
func numTrees(n int) int {
    if n == 0 {
        return 0
    }
    arr := combination(2*n)
    return arr[2*n][n]/(n+1)
}
func combination(n int) [][]int{
    var c [][]int = make([][]int,n+1)
    for i := 0 ; i <= n ; i++ {
        c[i] = make([]int,n+1)
    }
    for i := 0 ; i <= n ; i++{
        c[i][0] = 1
    }
    for i := 1 ; i <= n ; i++{
        for j := 1 ; j <= i ; j++ {
            c[i][j] = c[i-1][j-1]+c[i-1][j]
        }
    }
    return c
}