#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

/*
���⣺
�ж�һ�������е����������Ƿ����ظ��ģ�����з����棬���򷵻ؼ١�
˼·��
�����뵽�ķ������ǰ���������Ȼ�����һ�飬�ж��Ƿ������ڵ�������ȵ�����
*/
class Solution {
public:
	bool containsDuplicate(vector<int>& nums) {
		bool bRet = false;
		sort(nums.begin(), nums.end());
		for (int i = 1; i < nums.size(); ++i)
		{
			if (nums[i] == nums[i - 1])
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