func search(nums []int, target int) int {
    var hash = make(map[int]int)
    for _,num := range nums {
        hash[num]++
    }
    return hash[target]
}