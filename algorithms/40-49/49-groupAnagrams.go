//先将每个字符串按照ASCII码进行排序，保存到hash表中
func groupAnagrams(strs []string) [][]string {
    if strs == nil || len(strs) == 0 {
        return [][]string{}
    }
    rMap := make(map[string][]string, 0)
    for _, v := range strs {
        nstr := v
        nstr = aSort(nstr)
        if _, ok := rMap[nstr]; !ok {
            rMap[nstr] = []string{v}
        } else {
            rMap[nstr] = append(rMap[nstr], v)
        }
    }
    res := make([][]string, 0)
    for _, v := range rMap {
        a := make([]string, 0)
        for _, x := range v {
            a = append(a, x)
        }
        res = append(res, a)
    }
    return res
}

func aSort(s string) string{
    x := []byte(s)
    sort.Slice(x, func(i, j int) bool {return x[i] < x[j]})
    return string(x)
}