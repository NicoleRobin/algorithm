package main

import (
	"fmt"
	"sort"
)

/*
单调栈
*/
func minOperations(nums []int) int {
	var res int
	if len(nums) == 0 {
		return 0
	}

	if nums[0] != 0 {
		res = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			res++
		}
	}

	return res
}

func minOperations2(nums []int) int {
	zeroIndexMap := make(map[int]bool)
	numIndexListMap := make(map[int][]int)
	var uniqNumList []int
	for i, num := range nums {
		if num == 0 {
			zeroIndexMap[i] = true
		} else {
			if _, ok := numIndexListMap[num]; !ok {
				numIndexListMap[num] = []int{i}
				uniqNumList = append(uniqNumList, num)
			} else {
				numIndexListMap[num] = append(numIndexListMap[num], i)
			}
		}
	}

	var res int
	sort.Ints(uniqNumList)
	fmt.Println("uniqNumList:", uniqNumList)
	for _, minNum := range uniqNumList {
		var subArrayStart bool
		for i, num := range nums {
			if num == minNum {
				nums[i] = 0
				if !subArrayStart {
					res++
					subArrayStart = true
				}
			}
			if num == 0 {
				subArrayStart = false
			}
		}
	}

	return res
}

func minOperations1(nums []int) int {
	zeroIndexMap := make(map[int]bool)
	numIndexListMap := make(map[int][]int)
	var uniqNumList []int
	for i, num := range nums {
		if num == 0 {
			zeroIndexMap[i] = true
		} else {
			if _, ok := numIndexListMap[num]; !ok {
				numIndexListMap[num] = []int{i}
				uniqNumList = append(uniqNumList, num)
			} else {
				numIndexListMap[num] = append(numIndexListMap[num], i)
			}
		}
	}

	var res int
	sort.Ints(uniqNumList)
	fmt.Println("uniqNumList:", uniqNumList)
	for _, num := range uniqNumList {
		numIndexList := numIndexListMap[num]
		fmt.Println("numIndexList:", numIndexList)
		fmt.Println("zeroIndexMap:", zeroIndexMap)
		if len(numIndexList) == 1 {
			res++
			zeroIndexMap[numIndexList[0]] = true
		} else {
			res++
			zeroIndexMap[numIndexList[0]] = true
			for j := 1; j < len(numIndexList); j++ {
				zeroIndexMap[numIndexList[j]] = true
				for k := numIndexList[j-1] + 1; k < numIndexList[j]; k++ {
					if _, ok := zeroIndexMap[k]; ok {
						res++
						break
					}
				}
			}
		}
	}

	return res
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{1, 2, 1, 2, 1, 2},
			expected: 4,
		},
		{
			nums:     []int{0, 36, 0, 348, 347, 508, 62, 448, 568, 549, 0, 170, 0, 450, 68, 442, 16, 0, 500, 179, 462, 482, 321, 492, 0, 164, 443, 455, 495, 535, 67, 215, 35, 137, 0, 410, 282, 0, 464, 468, 281, 418, 192, 237, 123, 403, 115, 552, 0, 362, 311, 44, 501, 505, 448, 159, 0, 110, 178, 396, 166, 558, 581, 417, 591, 0, 281, 532, 498, 133, 499, 161, 370, 167, 530, 35, 29, 75, 485, 107, 351, 239, 166, 126, 574, 109, 505, 516, 423, 308, 227, 192, 107, 526, 19, 255, 0, 218, 65, 106, 425, 505, 97, 86, 297, 0, 136, 227, 0, 305, 360, 575, 243, 32, 245, 143, 181, 206, 0, 303, 85, 222, 544, 357, 482, 574, 341, 438, 489, 102, 444, 578, 139, 552, 75, 0, 182, 132, 367, 454, 0, 534, 318, 359, 423, 541, 351, 31, 82, 394, 0, 254, 211, 216, 341, 278, 148, 434, 123, 8, 493, 525, 466, 84, 0, 322, 89, 529, 302, 127, 387, 145, 564, 470, 289, 486, 246, 509, 463, 78, 529, 514, 178, 297, 377, 238, 134, 138, 89, 573, 0, 328, 341, 160, 232, 405, 568, 40, 105, 0, 0, 419, 225, 511, 359, 239, 0, 416, 0, 415, 170, 370, 418, 470, 323, 395, 349, 575, 491, 348, 568, 49, 438, 107, 248, 0, 409, 549, 413, 352, 143, 440, 187, 143, 426, 0, 148, 0, 570, 247, 134, 0, 152, 303, 132, 119, 23, 0, 291, 0, 229, 56, 299, 550, 547, 44, 259, 77, 375, 115, 420, 0, 123, 181, 62, 9, 236, 0, 0, 4, 496, 0, 120, 0, 254, 268, 337, 0, 0, 88, 373, 530, 401, 298, 536, 4, 140, 591, 551, 574, 470, 451, 169, 51, 0, 447, 509, 478, 255, 174, 319, 159, 376, 256, 293, 109, 410, 151, 0, 120, 442, 192, 213, 566, 74, 452, 13, 568, 500, 4, 356, 88, 325, 0, 538, 221, 313, 478, 202, 318, 459, 483, 233, 5, 167, 478, 31, 478, 449, 285, 79, 500, 279, 586, 119, 317, 0, 209, 335, 221, 1, 285, 517, 60, 245, 509, 548, 0, 234, 559, 33, 492, 319, 516, 271, 417, 243, 58, 345, 350, 333, 177, 493, 342, 168, 336, 195, 174, 418, 460, 541, 553, 405, 0, 0, 577, 59, 575, 350, 11, 44, 507, 513, 118, 158, 182, 420, 395, 0, 211, 387, 225, 42, 0, 48, 141, 0, 58, 0, 12, 273, 208, 569, 544, 0, 336, 562, 127, 445, 494, 549, 550, 571, 211, 86, 364, 262, 541, 181, 537, 266, 101, 214, 102, 274, 228, 391, 134, 181, 150, 378, 243, 295, 22, 351, 484, 154, 580, 562, 321, 69, 34, 515, 317, 0, 483, 547, 87, 342, 275, 514, 339, 357, 149, 349, 92, 381, 533, 105, 401, 462, 143, 105, 0, 527, 238, 66, 473, 64, 211, 0, 583, 213, 373, 164, 225, 211, 231, 474, 196, 387, 354, 490, 315, 36, 307, 30, 229, 489, 118, 106, 182, 0, 454, 158, 102, 124, 382, 554, 296, 431, 18, 535, 0, 37, 546, 256, 389, 471, 439, 48, 422, 483, 490, 261, 588, 0, 340, 160, 549, 139, 336, 150, 343, 280, 545, 0, 257, 19, 98, 559, 357, 458, 0, 0, 300, 462, 93, 375, 145, 313, 11, 495, 323, 0, 142, 584, 274, 382, 0, 400, 0, 199, 526, 0, 0, 0, 91, 208, 0, 211, 174, 188, 353, 274, 268, 309, 91, 286, 287, 0, 402, 477, 25, 173, 371, 303, 0, 112},
			expected: 519,
		},
	}

	for i, testCase := range testCases {
		res := minOperations(testCase.nums)
		fmt.Println("i:", i, "res:", res, "expected:", testCase.expected)
	}
}
