//将原来的数组分为三部分，第一部分是前面和新区间没有交集的，第二部分是和新区间有交集，
//第三部分是后面和新区间没有交集
func insert(intervals [][]int, newInterval []int) [][]int {
    var res [][]int
    k := 0
    for k < len(intervals) && intervals[k][1] < newInterval[0] {
        res = append(res,intervals[k])
        k++
    }
    if k < len(intervals) {
        newInterval[0] = min(newInterval[0],intervals[k][0])
        for k < len(intervals) && intervals[k][0] <= newInterval[1] {
            newInterval[1] = max(newInterval[1],intervals[k][1])
            k++
        }
    }
    res = append(res,newInterval)
     for k < len(intervals) && intervals[k][0] > newInterval[1] {
        res = append(res,intervals[k])
        k++
    }
    return res
}
func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}
func min(a,b int) int{
    if a < b {
        return a
    }
    return b
}