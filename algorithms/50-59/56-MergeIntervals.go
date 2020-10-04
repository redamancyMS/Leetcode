/*排序+合并
1.先对区间首部从小到大排序。
2.遍历、合并重合区间：
如果cs的尾部小于当前区间的头，即不相交，添加到结果集，并将尾部指到当前区间的尾部。
如果cs的尾部小于等于当前区间的尾，即有重合区域，合并： 将当前区间尾赋到cs的尾。

*/

func merge(intervals [][]int) [][]int {
    var res [][]int 
    if len(intervals) == 0 {
        return res
    }
     sort.Slice(intervals,func(i,j int)bool{
         return intervals[i][0] < intervals[j][0]
     })
     l := intervals[0][0]
     r := intervals[0][1]
     for i := 1 ; i < len(intervals) ; i++ {
         if intervals[i][0] > r {
             res = append(res,[]int{l,r})
             l = intervals[i][0]
             r = intervals[i][1]
         }else{
             r = max(r,intervals[i][1])
         }
     }
     res = append(res,[]int{l,r})
     return res
}
func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}