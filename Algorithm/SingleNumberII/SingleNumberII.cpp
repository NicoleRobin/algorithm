#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
���⣺
����һ�����֣�����һ�����֣��������ֶ����������Σ��ҳ�������֡�
�涨ʱ��Ϊ���Ը��Ӷ�
˼·��
���ȶ������������Ȼ��˳���жϣ�������������3����ȵ�������Ҫ�ҵġ�
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
˼·����
��Ȼʹ��λ���㣬��һ�����鱣����������Ӧλ��1�ĸ�������󽫸ø�����3���࣬����ǣ���ʾ����һ�ε�����λҲ��1��
���Ϊ0����ʾ����һ�ε�����λҲ��0.
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