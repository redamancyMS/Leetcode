func subsetsWithDup(nums []int) [][]int {
   var res [][]int
   n := len(nums)
   if n == 0 {
       return res
   }
   sort.Ints(nums)
   var path []int
   var dfs func([]int,int)
   dfs = func(num []int,u int){
       if u == len(num){
           res = append(res,append([]int(nil),path...))
           return
       }
       k := u+1
       for k < len(num) && num[k] == num[u]{
           k++
       }
       for i := 0 ; i <= k-u ; i++{
           dfs(num,k)
           path = append(path,num[u])
       }
       for i := 0 ; i <= k-u ; i++{
           path = path[:len(path)-1]
       } 
   }
   dfs(nums,0)
   return res
   
}