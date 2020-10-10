//先补位，再相加
func addBinary(a string, b string) string {
    m := len(a)
    n := len(b)
    for m < n {
        a = "0" + a
        m++
    }
    for n < m {
        b = "0" + b
        n++
    }
    var res string
    length := len(a)
    temp := 0
    for i := 0 ; i < length ; i++ {
        byteA,_ := strconv.Atoi(a[length-i-1:length-i])
        byteB,_ := strconv.Atoi(b[length-i-1:length-i])

        sum := temp + byteA + byteB
        if sum == 3 {
            temp = 1
            res = strconv.Itoa(1) + res
        }else if sum == 2 {
            temp = 1
            res = strconv.Itoa(0) + res
        }else if sum == 1 {
            temp = 0
            res = strconv.Itoa(1) + res
        }else {
            temp = 0
            res = strconv.Itoa(0) + res
        }
    }
    if temp != 0 {
        res = strconv.Itoa(1) + res
    }
    return res
    
}