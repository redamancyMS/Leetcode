func firstMissingPositive(nums []int) int {
    
    temp := make(map[int]int)
    for index,value := range nums {
        temp[value] = index
    }
    res := 1
    i := 0
    for i < len(nums){
        if _,ok := temp[res];ok {
            res++
        }
        i ++
    }
   
    return res
    
   
}

//时间复杂度为O(n)且只使用常数级别的额外空间

func firstMissingPositive(nums []int) int {
    
    for i := 0 ; i <len(nums) ; i++ {
        for nums[i] > 0 && nums[i] < len(nums) &&nums[i] != nums[nums[i]-1] {
            swap(&nums[i],&nums[nums[i]-1])
        }
    }
    for i := 0 ; i < len(nums) ; i ++ {
        if nums[i] != i + 1 {
            return i+1
        }
    }
    return len(nums)+1
   
}

func swap(a *int,b *int){
    *a,*b = *b,*a
}