package main

import (
	"fmt"
)

/*
思考：
1、根据题目要求，其实就是要选择连续的子数组，且子数组中的水果种类不能超过2个，求能摘到的最多水果个数；

思路：滑动窗口
1、left,right指针，初始均为0
2、向左移动right指针，并记录滑动窗口之间的值的个数
3、如果left、right之间的值种类大于2，则向左移动left指针，直至水果种类小于等于2
*/
func totalFruit(fruits []int) int {
	var res int
	var left int
	fruitMap := map[int]int{}
	for right := 0; right < len(fruits); right++ {
		fruitMap[fruits[right]]++
		// move left until len(fruitMap) <= 2
		for len(fruitMap) > 2 {
			// fmt.Println("fruitMap:", fruitMap)
			fruitMap[fruits[left]]--
			if fruitMap[fruits[left]] == 0 {
				delete(fruitMap, fruits[left])
			}
			left++
		}
		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}

func main() {
	testCases := []struct {
		fruits   []int
		expected int
	}{
		/*
			{
				fruits:   []int{1, 2, 1},
				expected: 3,
			},
			{
				fruits:   []int{0, 1, 2, 2},
				expected: 3,
			},
		*/
		{
			fruits:   []int{1, 2, 3, 2, 2},
			expected: 4,
		},
	}
	for i, tc := range testCases {
		res := totalFruit(tc.fruits)
		if res != tc.expected {
			fmt.Printf("#%d, expected:%d, res:%d\n", i, tc.expected, res)
		} else {
			fmt.Printf("#%d, pass\n", i)
		}
	}
}
