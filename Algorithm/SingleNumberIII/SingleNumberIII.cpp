#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
���⣺
����һ�����֣�������������ֻ����һ�Σ��������ֶ����������Σ��ҳ����������֡�
�涨ʱ��Ϊ���Ը��Ӷ�
˼·��
�ȶ���������
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