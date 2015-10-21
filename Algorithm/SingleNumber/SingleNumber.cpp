#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
���⣺
����һ�����֣�����һ�����ֳ���һ�Σ��������ֶ����������Σ��ҳ�������֡�
�涨ʱ��Ϊ���Ը��Ӷ�
˼·��
���ȶ������������Ȼ��˳���жϣ�������������2����ȵ�������Ҫ�ҵġ�
*/
class Solution {
public:
	int singleNumber(vector<int>& nums) {
		sort(nums.begin(), nums.end());
		int i;
		for (i = 1; i < nums.size(); i = i + 2)
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
˼·2��
ʹ������κ�����0��򣬵����������Լ���򣬵���0.
*/
class Solution {
public:
	int singleNumber(vector<int>& nums) {
		int nRet = 0;
		for (int i = 0; i < nums.size(); ++i)
		{
			nRet ^= nums[i];
		}
		return nRet;
	}
};


int main(void)
{


	return 0;
}