#include <iostream>
#include <string>

using namespace std;

/*
˼·��
��ת����
���Ⱦ���Ҫ�뵽������ʵ�����ȳ���-1������������Ȼ���ٳ���-1תΪ������
���������Ļ������Ƕ�10���࣬Ȼ�����10��һֱ������Ϊ0.
*/
class Solution {
public:
	int reverse(int x) {
		bool bNegative = false;
		if (x < 0)
		{
			bNegative = true;
			x *= -1;
		}

		long long lRet = 0;
		int i = 0;
		while (x)
		{
			lRet = lRet * 10 + x % 10;

			x /= 10;
		}

		if (bNegative)
		{
			lRet *= -1;
		}

		if (lRet > 2147483647 || lRet < -2147483648)
		{
			lRet = 0;
		}

		return (int)lRet;
	}
};

int main(void)
{
	Solution s;
	int n;
	while (true)
	{
		cin >> n;
		cout << s.reverse(n) << endl;
	}
	return 0;
}