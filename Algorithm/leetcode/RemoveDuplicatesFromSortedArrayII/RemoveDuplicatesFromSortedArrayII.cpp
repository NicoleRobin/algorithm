#include <iostream>
#include <vector>
using namespace std;

/*
���⣺
����һ�����飬���������е��������������Σ�����������ظ��������䳤�ȡ�
˼·��

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
˼·����
����������ظ����ֺ������Ӧ���ڵ�λ�ã����串��ԭ�������֡�
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
˼·����
������������������ͬ�����֣�ֻȡ����ͷ�����֣�����������
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