//满足以下这两个条件，生成的括号序列一定是有效的括号组合
//任意前缀里面左括号的数量大于等于右括号的数量
//左右括号数量相等
//任何情况下，只要左括号数量小于n就可以填进去，右括号小于n且任意前缀里面左括号的数量大于等于右括号的数量才可以填写右括号

func generateParenthesis(n int) []string {
     res := new([]string)
     dfs(n,0,0,"",res)
     return *res
}

func dfs(n int,ls int,rs int,sq string,res *[]string) {
    if ls == n && rs == n {
        *res = append(*res,sq)
    }else{
        if ls < n {
            dfs(n,ls+1,rs,sq + string('('),res)
        }
        if rs < n && ls > rs {
            dfs(n,ls,rs+1,sq + string(')'),res)
        }
    }
    
}