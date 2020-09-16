//可以和三数之和类比

func threeSumClosest(nums []int, target int) int {
    var min int = math.MinInt32
    if len(nums) < 3 {
        return min
    }
    sort.Ints(nums)
    var left,right,sum int
    min = abs(target-nums[0]-nums[1]-nums[2])
    sum = nums[0] + nums[1] + nums[2]
    for i := 0 ; i < len(nums) ; i++ {
        left = i + 1
        right = len(nums)-1
        for left < right {
            tmp := nums[i] + nums[left] + nums[right]
            if tmp == target {
                return target
            }
            if abs(target-tmp) < min {
                min = abs(target-tmp)
                sum = tmp
            }
            if tmp > target{
                right--
            }else if tmp < target {
                left++
            }
        }
    }
    return sum
}
func abs(x int) int {
    if x < 0 {
        return -1 * x
    }
    return x
}
