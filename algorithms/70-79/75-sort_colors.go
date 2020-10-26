//三指针，用j,i,k来对三种颜色进行划分，j,i从0开始++，k从len(colors)-1开始--,在遍历的过程中，
//保证从0到j-1都是0，从j到i-1都是1，从k+1到n-1都是2
func sortColors(colors []int) {
    i,j,k := 0,0,len(colors)-1
    for i <= k {
        if colors[i]==0 {
            swap(&colors[i],&colors[j])
            i++
            j++
        }else if colors[i]==2 {
            swap(&colors[i],&colors[k])
            k--
        }else{
            i++
        }
    }
}

func swap(a,b *int) {
    *a,*b = *b,*a
}