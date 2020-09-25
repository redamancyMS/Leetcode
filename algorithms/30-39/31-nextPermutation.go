/*找到第一个非降序的位置，
然后将其变成后面比它大的最小的那一个值，
最后将这个位置的后面的数按照升序进行排列，
就是最后的结果*/
func nextPermutation(nums []int)  {
    k := len(nums) - 1

    for k > 0 && nums[k-1] >= nums[k] {
        k--
    }
    if k <= 0 {
        //reverse
        for i,j := 0,len(nums)-1; i < j ; i,j=i+1,j-1 {
            nums[i],nums[j] = nums[j],nums[i]
        }
    }else{
        t := k
        for t < len(nums) && nums[t] > nums[k-1] {
            t++
        }
        swap(&nums[t-1],&nums[k-1])
        for i,j := k,len(nums)-1 ; i < j ; i,j = i+1,j-1 {
            nums[i],nums[j] = nums[j],nums[i]
        }
    }

}

func swap(a,b *int) {
    *a,*b = *b,*a
}