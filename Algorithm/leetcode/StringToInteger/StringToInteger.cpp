#include <iostream>
#include <string>

using namespace std;

/*
˼·��
��Ҫע��������¼��㣺
1���ַ����е�ǰ�ÿո���Ҫ������
2�������������ת��Ϊ���ֵ��ַ������������������Ϊ���������Ϊ0��
3������������ʱ�������Ϊ2147483647�������ʱ�������Ϊ-2147483648
*/
class Solution {
public:
	int myAtoi(string str) {
		size_t i = 0;
		bool bNegative = false;

		// ����ǰ�ÿո�
		while (str[i] == ' ')
		{
			i++;
		}

		// �ж���������
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