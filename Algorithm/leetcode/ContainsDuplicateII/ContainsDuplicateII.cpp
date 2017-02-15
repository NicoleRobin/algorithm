#include <iostream>
#include <vector>
#include <algorithm>
#include <stdlib.h>
using namespace std;

/*
���⣺
����һ������a��һ������k���ж��Ƿ�������������±�i,j����a[i] = a[j]��
��i j֮��Ĳ�����Ϊk��
˼·��
��ֱ�ӵ��뷨���Ǳ������������������顣�����ƽⳬʱ��
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
˼·����
�����򣬲���Ҫ����ֵ�Լ������±꣬Ȼ��������Ƿ�����ȵ������������֣�����У��ж��±���Ƿ�С�ڵ���k��
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