给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例 1:
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0

示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5

//用递归的方法去做
//求出两个数组中从小到大排列第k个数，k=(m+n)/2

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
     tot := len(nums1) + len(nums2)
     if tot % 2 == 0 {
         left := find(nums1,0,nums2,0,tot/2)
         right := find(nums1,0,nums2,0,tot/2 + 1)
         return (float64(left)+float64(right))/2.0
     }else{
         return float64(find(nums1,0,nums2,0,tot/2 +1)) 
     }
}

func find(nums1 []int,i int,nums2 []int,j int,k int) int {
    //为了方便计算，假定第一个数组是比较短的数组
    if len(nums1) - i > len(nums2) - j {
        return find(nums2,j,nums1,i,k)
    }
    //最后只剩下一个元素
    if k == 1 {
        //第一个数组为空
        if len(nums1) == i {
            return nums2[j]
        }else{
            return min(nums1[i],nums2[j])
        }
    }
    if len(nums1) == i {
        return nums2[j+k-1]
    }
    si := min(len(nums1),i+k/2 )
    sj := j+k-k/2
    if nums1[si-1] > nums2[sj-1]{
        return find(nums1,i,nums2,sj,k-(sj-j))
    }else {
        return find(nums1,si,nums2,j,k-(si-i))
    }
}

func min(a,b int) int {
    if a < b {
        return a 
    }
    return b
}
