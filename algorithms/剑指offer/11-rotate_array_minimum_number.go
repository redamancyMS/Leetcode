func minArray(numbers []int) int {
    //sort.Ints(numbers)
    //return numbers[0]

    // i := 1 
    // for  i < len(numbers) && numbers[i] >= numbers[i-1]{
    //     i++
    // }
    // if i == len(numbers) {
    //     return numbers[0]
    // }
    // return numbers[i]

    l := 0
    r := len(numbers) - 1
    for l < r {
        mid := (l+r) >> 1
       if numbers[mid] < numbers[r]{
           r = mid
       }else if numbers[mid] > numbers[r]{
           l = mid+1
       }else{
           r--
       }  
    }
    return numbers[l]
}