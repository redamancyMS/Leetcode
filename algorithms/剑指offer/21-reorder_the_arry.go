//采用双指针的办法，有点类似快排
func exchange(nums []int) []int {
    i,j := 0,len(nums)-1

    for i <= j {
        for i <= len(nums)-1 && nums[i] % 2 == 1 {
            i++
        }
        for j >= 0 && nums[j] % 2 == 0 {
            j--
        }
        if i <=j {
            exchangePos(&nums[i],&nums[j])
            i++
            j--
        }
    }
    return nums
}

func exchangePos(a,b *int){
    *a,*b = *b,*a
}