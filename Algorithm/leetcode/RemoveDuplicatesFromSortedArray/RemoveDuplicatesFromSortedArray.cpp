#include <iostream>
#include <algorithm>
#include <vector>
#include <iterator>
using namespace std;

/*
���⣺
����һ����������飬ȥ�������ظ������֣����������ĳ��ȡ�
˼·��
�������飬ȥ���ظ�����
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
˼·����
ʹ��STL�е�unique
*/
class Solution {
public:
	int removeDuplicates(vector<int>& nums) {
		return distance(nums.begin(), unique(nums.begin(), nums.end()));
	}
};

/*
˼·����
��һ��������¼�����ظ���ÿ��Ԫ�ض�Ӧ��λ�ã��������飬��Ԫ�طŵ���ȷ��λ�á�
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