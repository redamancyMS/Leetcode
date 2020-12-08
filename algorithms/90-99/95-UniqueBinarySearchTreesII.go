/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//由1到n形成的二叉搜索树的个数是固定的，是卡特兰数C(下:2n,上:n)/(n+1)
//从1-n形成的二叉搜索树，因为二叉搜索树的中序遍历是递增的，所以在遍历的时候每棵子树都是连续的一段
func generateTrees(n int) []*TreeNode {
     if n == 0 {
         return nil
     }
     var dfs func(int,int) []*TreeNode
    dfs = func(l int,r int) []*TreeNode{
        if l > r { //说明当前区间内没有节点，子树为空
            return []*TreeNode{nil}
        }
        var res []*TreeNode
        for i:= l ; i <= r ; i++ {
            left := dfs(l,i-1)
            right := dfs(i+1,r)
            for _,l := range left{
                for _,r := range right{
                    root := &TreeNode{i,nil,nil}
                    root.Left = l
                    root.Right = r
                    res = append(res,root)
                }
            }
        }
        return res

    }
    return dfs(1,n)
}