#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
题意：
给定一组数字，除了一个数字，其他数字都出现了三次，找出这个数字。
规定时间为线性复杂度
思路：
首先对数组进行排序，然后顺序判断，遇到不是连续3个相等的数就是要找的。
*/
class Solution {
public:
	int singleNumber(vector<int>& nums) {
		sort(nums.begin(), nums.end());
		int i;
		for (i = 1; i < nums.size(); i = i + 3)
		{
			if (nums[i] != nums[i - 1])
			{
				break;
			}
		}
		return nums[i - 1];
	}
};

/*
思路二：
仍然使用位运算，用一个数组保存所有数对应位上1的个数，最后将该个数对3求余，结果是，表示出现一次的数该位也是1，
如果为0，表示出现一次的数该位也是0.
*/
class Solution
{
public:
	int singleNumber(vector<int>& nums) {
		int nWidth = sizeof(int) * 8;
		int bitNum[nWidth] = {0};
		int nRet = 0;
		for (int i = 0; i < nWidth; ++i)
		{
			for (int j = 0; j < nums.size(); ++j)
			{
				bitNum[i] += (nums[j] >> i) & 1;
			}
			nRet |= (bitNum[i] % 3) << i;
		}
		return nRet;
	}
};

int main(void)
{
	

	return 0;
}