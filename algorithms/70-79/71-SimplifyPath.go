func simplifyPath(path string) string {
    //利用栈的思想
    var res string
    if len(path) == 0 {
        return res
    }
    n := len(path)
    if path[n-1] != '/' {
        path += "/"
    }
    var name string // 用来暂时存储前面的路径名
    for _,c := range path {
        if c != '/' {
            name += string(c)
        }else{
            //上一级目录
            if name == ".."  {
                for len(res) > 0 && res[len(res)-1] !='/' {
                    res = res[:len(res)-1]
                }
                if len(res) > 0 {
                    res = res[:len(res)-1]
                }
            }else if name != "." && name != "" {
                res += "/" + name
            }
            name = ""
        }
    }
    if len(res) == 0 {
        res = "/"
    }
    return res
}