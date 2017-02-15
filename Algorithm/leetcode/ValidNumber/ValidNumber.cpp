#include <iostream>
#include <string>

using namespace std;

/*
思路：
考虑用状态机来实现，不同的类型对应不同的状态。
但是按照自己的思路实现了之后却没有AC，因为.1按我的思路它不是数字所以返回false，但是答案却说它是true。
我就奇怪了.1怎么就是数了？

经过20次的submit终于AC了。
*/
enum ENState
{
	INIT,		// 初始
	START,		// 开始
	INTEGER,	// 整型
	FLOAT,		// 小数
	SCIENCE,	// 科学
	INDEX,		// 指数
	END,		// 结尾
	ERROR		// 错误
};
class Solution {
public:
	bool isNumber(string s) {
		ENState state = INIT;
		bool bIsNumeric = false;
		for (size_t i = 0; i < s.length(); ++i)
		{
			switch (state)
			{
			case INIT:
				if (s[i] >= 48 && s[i] <= 57)
				{ // 整数
					state = INTEGER;
					bIsNumeric = true;
				}
				else if (s[i] == '+' || s[i] == '-')
				{ // 开始
					state = START;
					bIsNumeric = false;
				}
				else if (s[i] == '.')
				{ // 浮点型
					state = FLOAT;
					bIsNumeric = false;
				}
				else if (s[i] == ' ' || s[i] == '\t' || s[i] == '\n')
				{ // 仍是初始
					state = INIT;
					bIsNumeric = false;
				}
				else
				{ // 错误
					state = ERROR;
					bIsNumeric = false;
				}
				break;
			case START:
				if (s[i] >= 48 && s[i] <= 57)
				{
					state = INTEGER;
					bIsNumeric = true;
				}
				else if (s[i] == '.')
				{
					state = FLOAT;
					bIsNumeric = false;
				}
				else
				{
					state = ERROR;
				}
				break;
			case INTEGER:
				if (s[i] >= 48 && s[i] <= 57)
				{
					state = INTEGER;
					bIsNumeric = true;
				}
				else if (s[i] == '.')
				{
					state = FLOAT;
					bIsNumeric = true;
				}
				else if (s[i] == 'e')
				{
					state = SCIENCE;
					bIsNumeric = false;
				}
				else if (s[i] == ' ' || s[i] == '\t' || s[i] == '\n')
				{
					state = END;
				}
				else
				{
					state = ERROR;
					bIsNumeric = false;
					i = s.length();
				}
				break;
			case FLOAT:
				if (s[i] >= 48 && s[i] <= 57)
				{
					state = FLOAT;
					bIsNumeric = true;
				}
				else if (bIsNumeric && s[i] == 'e')
				{
					state = SCIENCE;
					bIsNumeric = false;
				}
				else if (s[i] == ' ' || s[i] == '\t' || s[i] == '\n')
				{
					state = END;
				}
				else
				{
					state = ERROR;
					bIsNumeric = false;
					i = s.length();
				}
				break;
			case SCIENCE:
				if (s[i] >= 48 && s[i] <= 57)
				{
					state = INDEX;
					bIsNumeric = true;
				}
				else if (s[i] == '+' || s[i] == '-')
				{
					state = INDEX;
					bIsNumeric = false;
				}
				else if (s[i] == ' ' || s[i] == '\t' || s[i] == '\t')
				{
					state = END;
					bIsNumeric = false;
				}
				else
				{
					state = ERROR;
					bIsNumeric = false;
				}
				break;
			case INDEX:
				if (s[i] >= 48 && s[i] <= 57)
				{
					state = INDEX;
					bIsNumeric = true;
				}
				else if (s[i] == ' ' || s[i] == '\t' || s[i] == '\t')
				{
					state = END;
				}
				else
				{
					state = ERROR;
					bIsNumeric = false;
				}
				break;
			case END:
				if (s[i] == ' ' || s[i] == '\t' || s[i] == '\t')
				{

				}
				else
				{
					state = ERROR;
					bIsNumeric = false;
					i = s.length();
				}
				break;
			case ERROR:
				bIsNumeric = false;
				i = s.length();
				break;
			default:
				break;
			}
		}

		return bIsNumeric;
	}
};

int main(void)
{
	Solution s;
	string str;
	while (true)
	{
		cin >> str;
		cout << s.isNumber(str) << endl;
	}

	return 0;
}