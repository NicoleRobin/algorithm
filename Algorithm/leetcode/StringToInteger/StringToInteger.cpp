#include <iostream>
#include <string>

using namespace std;

/*
思路：
需要注意的有以下几点：
1、字符串中的前置空格需要消除。
2、如果遇到不能转换为数字的字符，则结束，而不是认为出错将结果置为0。
3、当结果上溢出时将结果置为2147483647，下溢出时将结果置为-2147483648
*/
class Solution {
public:
	int myAtoi(string str) {
		size_t i = 0;
		bool bNegative = false;

		// 消除前置空格
		while (str[i] == ' ')
		{
			i++;
		}

		// 判断整数符号
		if (str[i] == '-')
		{
			bNegative = true;
			i++;
		}
		else if (str[i] == '+')
		{
			i++;
		}

		long long lRet = 0;
		for (; i < str.length(); ++i)
		{
			if (str[i] < 48 || str[i] > 57)
			{
				break;
			}
			lRet = lRet * 10 + (str[i] - 48);
			if (lRet > 2147483647)
			{
				break;
			}
		}

		if (bNegative)
		{
			lRet *= -1;
		}

		if (lRet > 2147483647 || lRet < -2147483648)
		{
			if (bNegative)
			{
				lRet = -2147483648;
			}
			else
			{
				lRet = 2147483647;
			}
		}
		return (int)lRet;
	}
};

int main(void)
{
	Solution s;
	string str;
	while (true)
	{
		cin >> str;
		cout << s.myAtoi(str) << endl;
	}

	return 0;
}