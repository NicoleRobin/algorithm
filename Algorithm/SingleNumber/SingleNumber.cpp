#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
题意：
给定一组数字，除了一个数字，其他数字都出现了两次，找出这个数字。
规定时间为线性复杂度
思路：
首先对数组进行排序，然后顺序判断，遇到不是连续2个相等的数就是要找的。
*/
class Solution {
public:
	int singleNumber(vector<int>& nums) {
		sort(nums.begin(), nums.end());
		int i;
		for (i = 1; i < nums.size(); i = i + 2)
		{
			if (nums[i] != nums[i - 1])
			{
				break;
			}
		}
		return nums[i - 1];
	}
};

int main(void)
{


	return 0;
}