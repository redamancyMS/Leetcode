给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案,但是，数组中同一个元素不能使用两遍。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]


//用map来做比较合适，在遍历数组的过程中，每次都在map中查找是否存在target-nums[i]，如果存在，直接返回二者的下标。

func twoSum(nums []int, target int) []int {
   temp := make(map[int]int)
    res := []int{}
   for i := 0 ; i < len(nums) ; i++ {
       num := target-nums[i]
       if _,ok := temp[num] ; ok {
           res = append(res,i)
           res = append(res,temp[num])
       }else{
           temp[nums[i]] = i
       }
   }
   return res
}
