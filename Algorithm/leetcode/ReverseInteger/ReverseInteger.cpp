#include <iostream>
#include <string>

using namespace std;

/*
思路：
翻转整数
首先就需要想到负数其实可以先乘以-1，按正数处理，然后再乘以-1转为负数。
对于正数的话，就是对10求余，然后除以10，一直到该数为0.
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