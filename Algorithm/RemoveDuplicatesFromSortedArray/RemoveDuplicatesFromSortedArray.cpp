#include <iostream>
#include <algorithm>
#include <vector>
#include <iterator>
using namespace std;

/*
题意：
给定一个有序的数组，去掉其中重复的数字，并返回最后的长度。
思路：
遍历数组，去掉重复数字
*/
class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		if (nums.empty())
		{
			return 0;
		}
		else
		{
			for (int i = 0; i < nums.size() - 1;)
			{
				if (nums[i] == nums[i + 1])
				{
					nums.erase(nums.begin() + i);
				}
				else
				{
					++i;
				}
			}
			return nums.size();
		}
	}
};

/*
思路二：
使用STL中的unique
*/
class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		return distance(nums.begin(), unique(nums.begin(), nums.end()));
	}
};

/*
思路三：
用一个变量记录消除重复后每个元素对应的位置，遍历数组，将元素放到正确的位置。
*/
class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		int nLen = 0;
		if (nums.size() == 0)
		{
			return nLen;
		}

		for (int i = 1; i < nums.size(); ++i)
		{
			if (nums[nLen] != nums[i])
			{
				nums[++nLen] = nums[i];
			}
		}
		return nLen;
	}
};

int main(void)
{
	Solution s;
	int n;
	vector<int> vecN;
	while (cin >> n)
	{
		vecN.push_back(n);
	}
	cout << s.removeDuplicates(vecN) << endl;
	copy(vecN.begin(), vecN.end(), ostream_iterator<int>(cout, ","));
	return 0;
}