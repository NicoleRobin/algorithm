#include <iostream>
#include <vector>
#include <algorithm>
#include <fstream>
#include <iterator>

using namespace std;

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
	vector<int> twoSum(vector<int>& nums, int target)
	{
		vector<int> res;
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

		int i, j;
		int sum;
		for (i = 0, j = nodes.size() - 1; i < j;)
		{
			sum = nodes[i].num + nodes[j].num;
			if (sum == target)
			{ // equal
				if (nodes[i].pos < nodes[j].pos)
				{
					res.push_back(nodes[i].pos);
					res.push_back(nodes[j].pos);
				}
				else
				{
					res.push_back(nodes[j].pos);
					res.push_back(nodes[i].pos);
				}
				break;
			}
			else if (sum > target)
			{ // grater than
				j--;
			}
			else
			{ // less than
				i++;
			}
		}
		return res;
	}
};

int main()
{
	Solution towsum;
	vector<int> nums;
	vector<int> res;
	nums.push_back(2);
	nums.push_back(3);
	nums.push_back(4);

	res = towsum.twoSum(nums, 6);

	copy(res.begin(), res.end(), ostream_iterator<int>(cout, ","));
}