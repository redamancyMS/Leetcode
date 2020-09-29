// 解法思想： 两墙夹一水，在找到右墙后逐步回退，与左边的每一个墙做结算
// 使用栈来处理
func trap1(height []int) int {
	sum := 0
	stack := make([]int, 0)

	current := 0

	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			h := height[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			min := min(height[stack[len(stack)-1]], height[current])
			sum += distance * (min - h)
		}

		stack = append(stack, current)
		current++

	}
	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
例子
    = 0 0 0 =    = 代表柱子
= 0 = = 0 = =    0 代表雨水
- - - - - - -    - 代表地面
0 1 2 3 4 5 6

 1.获取 0 位置的柱子高度,栈为空, 横坐标加入栈中,
      stack: 0 (此处是横坐标,不是高度)
 2.获取 1 位置的柱子高度, 不大于栈顶元素的高度, 横坐标加入栈中
     stack: 0 , 1
 3.获取 2 位置的柱子高度, 比栈顶元素的高度大, 说明中间可能有水, 开始计算(当前元素未入栈)
       1.将栈顶元素 1 出栈. h=0 ( h 可以理解成池塘底的高度 ), 出栈后, 栈中剩余 stack: 0 , 此时的栈顶元素就是 0 , 计算和此时栈顶元素(也就是墙)的积水量 (2-0-1)*(min(2,1)-0)=1, sum+=1 sum=1 (这里有点绕, 但是很精妙).
       2.此时栈已空, 循环结束, 将 2 入栈  
          stack: 2 ,
 4.获取 3 位置的柱子高度, 不大于栈顶元素的高度, 横坐标加入栈中
      stack: 2 , 3
 5.获取 4 位置的柱子高度, 不大于栈顶元素的高度, 横坐标加入栈中
       stack: 2 , 3 , 4
 6.获取 5 位置的柱子高度, 比栈顶元素的高度大, 说明中间可能有水, 开始计算(当前元素未入栈)
       将栈顶元素 4 出栈.h=0 出栈后, 栈中剩余 stack: 2 , 3 ,此时的栈顶元素就是 3 , 计算和此时栈顶元素(也就是墙)的积水量 (5-3-1)*(min(1,1)-0)=1, sum+=1 sum=2 ,
        stack: 2 , 3 , 5
 7.获取 6 位置的柱子高度, 比栈顶元素的高度大, 说明中间可能有水, 开始计算(当前元素未入栈)
    1. 将栈顶元素 5 出栈.h=1, 出栈后, 栈中剩余 stack: 2 , 3 ,此时的栈顶元素就是 3 , 计算和此时栈顶元素(也就是墙)的积水量 (6-3-1)*(min(1,2)-1)=0, sum+=0 sum=2 ,
    2. 栈顶元素 3 的高度比当前柱子的高度低,故继续循环. 将栈顶元素 3 出栈.h=1, 出栈后, 栈中剩余 stack: 2 ,此时的栈顶元素就是 2 , 计算和此时栈顶元素(也就是墙)的积水量 (6-2-1)*(min(2,2)-1)=3, sum+=3 sum=5 ,
    3. 当前位置 5 的柱子高度 不比栈顶元素的高度高, 循环结束, 将 6 入栈
        stack: 2 , 6
 8.height 数组遍历完成, 循环结束.
*/