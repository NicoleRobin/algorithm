#include <iostream>
using namespace std;

/*
���⣺
����һ���������������Ӧ�Ķ�������1�ĸ�����
˼·��
�Ը�����һ�������ȶ�2���࣬���Ƿ�Ϊ1��Ȼ���2�����ѭ��ֱ������Ϊ0.
*/
class Solution {
public:
	int hammingWeight(unsigned int n) {
		int nCount = 0;
		while (n > 0)
		{
			if (n % 2 == 1)
			{
				++nCount;
			}
			n /= 2;
		}
		return nCount;
	}
};

int main(void)
{
	unsigned int n;
	Solution s;
	while (true)
	{
		cin >> n;
		cout << s.hammingWeight(n) << endl;
	}

	return 0;
}