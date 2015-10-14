#include <iostream>
#include <string>

using namespace std;

/*
思路：
*/
enum ENState
{
	START,		// 初始
	INTEGER,	// 整型
	FLOAT,		// 小数
	SCIENCE,	// 科学计数法
	END,		// 结尾
	ERROR		// 错误
};
class Solution {
public:
	bool isNumber(string s) {
		ENState state = START;
		bool bIsNumeric = false;
		for (size_t i = 0; i < s.length(); ++i)
		{
			switch (state)
			{
			case START:
				if (s[i] >= 48 && s[i] <= 57)
				{
					state = INTEGER;
					bIsNumeric = true;
				}
				else if (s[i] == '+' || s[i] == '-' || s[i] == ' ')
				{

				}
				else
				{
					state = ERROR;
				}
				break;
			case INTEGER:
				if (s[i] >= 48 && s[i] <= 57)
				{

				}
				else if (s[i] == '.')
				{
					state = FLOAT;
					bIsNumeric = false;
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
					bIsNumeric = true;
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
					bIsNumeric = true;
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