#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

/*
题意：
判断一个数组中的所有数字是否有重复的，如果有返回真，否则返回假。
思路：
首先想到的方法就是把数组排序，然后遍历一遍，判断是否有相邻的两个相等的数。
*/
class Solution {
public:
	bool containsDuplicate(vector<int>& nums) {
		bool bRet = false;
		sort(nums.begin(), nums.end());
		for (int i = 1; i < nums.size(); ++i)
		{
			if (nums[i] == nums[i - 1])
			{
				bRet = true;
				break;
			}
		}
		return bRet;
	}
};

int main(void)
{


	return 0;
}