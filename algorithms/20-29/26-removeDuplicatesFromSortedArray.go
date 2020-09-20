//这个问题的重点需要关注：原地修改数组，并且不需要考虑数组中超出新长度后面的元素
func removeDuplicates(nums []int) int {
   if len(nums) == 0 {
       return 0
   }
   pre := nums[0]
   count := 1
   for _,cur := range nums {
       if cur != pre{
           nums[count] = cur
           count++
           pre = cur
       }
   }
   return count
}