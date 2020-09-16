var res []string 
var strs map[string]string = map[string]string{
    "2":"abc",
    "3":"def",
    "4":"ghi",
    "5":"jkl",
    "6":"mno",
    "7":"pqrs",
    "8":"tuv",
    "9":"wxyz"}
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    res = []string{}
    dfs(digits,0,"")
    return res
    
}
func dfs(digits string,u int , path string) {
    if u == len(digits) {
        res = append(res,path)
    }else{
        digit := string(digits[u])
        letters := strs[digit]
        for i:= 0 ; i < len(letters) ; i++{
            dfs(digits,u+1,path+string(letters[i]))
        }
    }
}