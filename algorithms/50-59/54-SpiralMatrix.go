func spiralOrder(matrix [][]int) (res []int) {

    if len(matrix) == 0 {
        return res
    }
    left := 0
    right := len(matrix[0])-1
    up := 0
    down := len(matrix)-1
   for true {
       for i :=left ; i <= right ; i++ {
           res = append(res,matrix[up][i])
       }
       up ++
       if up > down {
           break
       }
       for i := up ; i <= down ; i++ {
           res = append(res,matrix[i][right])
       }
       right--
       if right < left {
           break
       }
       for i:= right ; i >= left ; i-- {
           res = append(res,matrix[down][i])
       }
       down--
       if down < up {
           break
       }
       for i := down ; i >= up ; i-- {
           res = append(res,matrix[i][left])
       }
       left++
       if left > right {
           break
       }
   }
   return res
}