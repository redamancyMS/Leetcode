//采用竖式计算的方法，不过要在最后一步采取进位
func multiply(num1 string, num2 string) string {
    m := len(num1)
    n := len(num2)
    var arr1 []int
    var arr2 []int
    for i := m-1 ; i >= 0 ; i-- {
        arr1 = append(arr1,int(num1[i]-'0'))
    }
    for j := n-1 ; j >= 0 ; j-- {
        arr2 = append(arr2,int(num2[j]-'0'))
    }
    sum := make([]int,m+n)
    for i := 0 ; i < m ; i++ {
        for j := 0 ; j < n ; j++ {
            sum[i+j] += arr1[i]*arr2[j]
        }
    }
    for i,t := 0,0 ; i < len(sum) ; i++ {
        t += sum[i]
        sum[i] = t % 10
        t = t / 10
    }
    k := len(sum)-1
    for k > 0 && sum[k] == 0 {
        k--
    }
    var res string
    for k >= 0 {
        res += strconv.Itoa(sum[k])
        k--
    }
    return res
}