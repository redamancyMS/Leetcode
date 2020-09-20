func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    count := 0
    for _,num := range nums {
        if num != val {
            nums[count] = num
            count++
        }
    }
    return count
}