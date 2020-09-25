//找出分段点 
//题目要求使用logn的时间复杂度，所以要考虑二分
func search(nums []int, target int) int {
     if len(nums) == 0 {
         return -1
     }
     l := 0
     r := len(nums)-1
     for l < r {
         mid := l + r + 1 >> 1
         if nums[mid] >= nums[0] {
             l = mid
         }else{
             r = mid -1 
         }
     }
     if target >= nums[0] {
         l = 0
     }else {
         l = r + 1
         r = len(nums) - 1
     }

     for l < r {
         mid := (l+r)/2
         if nums[mid] >= target {
             r = mid
             
         }else{
            l = mid + 1
         }
     }
     if nums[r] == target {
         return r
     }
     return -1
}