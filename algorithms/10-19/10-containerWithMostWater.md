给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

![image](https://github.com/redamancyMS/Leetcode/blob/master/algorithms/images/container.png)

示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49

// 采用双指针的方法，从两边往中间寻找
func maxArea(height []int) int {
    if len(height) < 2 {
        return 0
    }
    i := 0
    j := len(height) -1 
    cap := (j-i) * min(height[i],height[j])
    for i < j {
        if height[i] < height[j] {
            i++
        }else if height[j] <= height[j] {
            j--
        }
        if i < j {
            cap = max(cap,(j-i)*min(height[i],height[j]))
        }
        
    }
    return cap
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b 
}
