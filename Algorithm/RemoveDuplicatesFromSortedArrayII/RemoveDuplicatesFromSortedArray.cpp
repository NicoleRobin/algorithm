#include <iostream>
#include <vector>
using namespace std;

/*
题意：
给定一个数组，允许数组中的数字最多出现两次，消除更多的重复，并求其长度。
思路：

*/
class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		if (nums.size() == 0)
		{
			return 0;
		}
		int nCount = 1;
		for (int i = 0; i < nums.size() - 1; ++i)
		{
			if (nums[i] == nums[i + 1])
			{
				nCount++;
				if (nCount > 2)
				{
					nums.erase(nums.begin() + i);
					i--;
				}
			}
			else
			{
				nCount = 1;
			}
		}

		return nums.size();
	}
};
/*
思路二：
计算出消除重复数字后的数字应该在的位置，将其覆盖原来的数字。
*/
class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		if (nums.size() <= 2)
		{
			return nums.size();
		}

		int nLen = 2;
		for (int i = 2; i < nums.size(); ++i)
		{
			if (nums[nLen - 2] != nums[i])
			{
				nums[nLen++] = nums[i];
			}
		}
		return nLen;
	}
};

/*
思路三：
对于任意多个相邻且相同的数字，只取其两头的数字，其他舍弃。
*/
class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		int nLen = 0;
		for (int i = 0; i < nums.size(); ++i)
		{
			if (i > 0 && i < nums.size() - 1 && nums[i - 1] == nums[i] && nums[i] == nums[i + 1])
			{
				continue;
			}
			nums[nLen++] = nums[i];
		}
		return nLen;
	}
};

int main(void)
{
	return 0;
}