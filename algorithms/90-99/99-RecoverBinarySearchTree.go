/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode)  {
    var first *TreeNode = nil //逆序对的第一个值
    var second *TreeNode
    var last *TreeNode = nil  //寻找前驱节点

    for root != nil {
        if root.Left == nil {
            if last != nil && last.Val > root.Val {
                if first == nil {
                    first = last
                    second = root
                }else{
                    second = root
                }
            }
            last = root
            root = root.Right
        }else{
            p := root.Left
            for p.Right!=nil && p.Right != root {
                p = p.Right
            }
            if p.Right == nil {
                p.Right = root
                root = root.Left
            }else{
                p.Right = nil
                if last != nil && last.Val > root.Val{
                    if first == nil{
                        first = last
                        second = root
                    }else{
                        second = root
                    } 
                }
                last = root
                root = root.Right
            }
        }
    }
    swap(&first.Val,&second.Val)
}
func swap(a,b *int) {
    *a,*b = *b,*a
}