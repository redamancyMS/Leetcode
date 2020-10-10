//剩下t个空格，一共有c个间隙，每个间隙分配t/c个空格，剩下t%c个空格依次从左往右每个间隙放一个空格
//情况：该一行只剩一个单词，左对齐；最后一行剩多个单词，左对齐，每个单词间一个空格；左右对齐
func fullJustify(words []string, maxWidth int) []string {
    var res []string
    for i := 0 ; i < len(words) ; i++ {
        j := i + 1
        length := len(words[i])
        for j < len(words) && length + 1 + len(words[j]) <=maxWidth{
            length += 1 + len(words[j])
            j++
        }
        var line string
        if j == len(words) || j == i+1 { //左对齐
            line += words[i]
            for k := i+1 ; k < j ; k++ {
                line += " " + words[k]
            }
            for len(line) < maxWidth {
                    line += " "
            }
        }else{ // 左右对齐
            c := j-i-1 //有多少个间隙
            t := maxWidth-length + c // 有多少空格
            line += words[i]
            k := 0
            for k < t % c {
                for num := 0 ; num < t/c +1 ; num ++{
                    line += " "
                }
                line += words[i+k+1]
                k++
            }
            for k < c {
                for num := 0 ; num < t/c ; num++ {
                    line += " "
                }
                line += words[i+k+1]
                k++
            }

        }
        res = append(res,line)
        i= j-1

    }
    return res
}