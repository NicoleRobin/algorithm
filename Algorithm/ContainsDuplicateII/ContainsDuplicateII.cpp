#include <iostream>
#include <vector>
#include <algorithm>
#include <stdlib.h>
using namespace std;

/*
题意：
给定一个数组a和一个整数k，判断是否存在这样两个下标i,j满足a[i] = a[j]，
且i j之间的差至多为k。
思路：
最直接的想法就是暴力搜索，遍历该数组。暴力破解超时。
*/
class Solution {
public:
	bool containsNearbyDuplicate(vector<int>& nums, int k) {
		bool bRet = false;
		for (int i = 0; i < nums.size(); ++i)
		{
			for (int j = i + 1; j <= i + k; ++j)
			{
				if (nums[i] == nums[j])
				{
					bRet = true;
					break;
				}
			}
		}
		return bRet;
	}
};

/*
思路二：
先排序，不过要保留值以及它的下标，然后遍历看是否有相等的两个相邻数字，如果有，判断下标差是否小于等于k。
*/
struct Node
{
	int num;
	int pos;
	bool operator<(Node node)
	{
		return num < node.num;
	}
};

class Solution
{
public:
	bool containsNearbyDuplicate(vector<int>& nums, int k) {
		vector<Node> nodes;
		Node node;
		vector<int>::iterator itIdx = nums.begin();
		for (; itIdx != nums.end(); ++itIdx)
		{
			node.num = *itIdx;
			node.pos = itIdx - nums.begin() + 1;
			nodes.push_back(node);
		}

		sort(nodes.begin(), nodes.end());

		int i;
		bool bRet = false;
		for (i = 1; i < nodes.size(); ++i)
		{
			if (nodes[i - 1].num == nodes[i].num && abs(nodes[i - 1].pos - nodes[i].pos) <= k)
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