func restoreIpAddresses(s string) []string {
    var ans []string
    var dfs func(string,int,int,string)
    //u表示当前搜到第几位，k表示当前搜到第几个数
    dfs = func(str string,u int,k int,path string){
        if u == len(s) {
            if k == 4 {
               s := string(path)
               c := []rune(s)
               c = c[:len(c)-1]
                ans = append(ans,string(c))
                return 
            }
        }
        if k == 4 {return}
        for i,t := u,0 ; i < len(s) ; i++ {
            if i > u && s[u] == '0'{break}//有前导0
            t = t*10 + int(s[i]-'0') 
            if t <= 255{
                dfs(s,i+1,k+1,path+strconv.Itoa(t)+".")
            }else{break}
        }
    }
    dfs(s,0,0,"")
    return ans
}