//双指针操作从后往前更新数组
func merge(nums1 []int, m int, nums2 []int, n int)  {
    a := m - 1
    b := n - 1
    c := m + n - 1
    for a >= 0 && b >= 0 {
        if nums1[a] > nums2[b] {
            nums1[c] = nums1[a]
            a--
            c--
        }else{
            nums1[c] = nums2[b]
            b--
            c--
        }
    }
    for b >= 0 {
        nums1[c] = nums2[b]
        b--
        c--
    }

}