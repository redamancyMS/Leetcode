//1、首先沿着对角线翻转，以对角线为轴将元素对称交换
//2、然后沿着中轴翻转，以中线为轴将元素对称交换
func rotate(matrix [][]int)  {
    n := len(matrix)
    for i := 0 ; i < n ; i++ {
        for j := 0 ; j < i ; j ++ {
            swap(&matrix[i][j],&matrix[j][i])
        }
    }

    for i := 0 ; i < n ; i ++ {
        for j,k := 0,n-1 ; j < k ; j,k = j + 1,k-1 {
            swap(&matrix[i][j],&matrix[i][k])
        }
    }

}

func swap(a,b *int) {
    *a,*b = *b,*a
}