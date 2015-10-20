#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
题意：
给定一组数字，除了两个数字只出现一次，其他数字都出现了两次，找出这两个数字。
规定时间为线性复杂度
思路：
先对数组排序，
*/
class Solution {
public:
	vector<int> singleNumber(vector<int>& nums) {
		sort(nums.begin(), nums.end());
		vector<int> vecRet;
		int i;
		for (i = 1; i < nums.size(); i = i + 2)
		{
			if (nums[i] != nums[i - 1])
			{
				vecRet.push_back(nums[i - 1]);
				if (vecRet.size() == 2)
				{
					break;
				}
				--i;
			}
		}
		return vecRet;
	}
};

int main(void)
{


	return 0;
}