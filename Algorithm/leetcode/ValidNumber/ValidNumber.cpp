#include <iostream>
#include <string>

using namespace std;

/*
˼·��
������״̬����ʵ�֣���ͬ�����Ͷ�Ӧ��ͬ��״̬��
���ǰ����Լ���˼·ʵ����֮��ȴû��AC����Ϊ.1���ҵ�˼·�������������Է���false�����Ǵ�ȴ˵����true��
�Ҿ������.1��ô�������ˣ�

����20�ε�submit����AC�ˡ�
*/
enum ENState
{
	INIT,		// ��ʼ
	START,		// ��ʼ
	INTEGER,	// ����
	FLOAT,		// С��
	SCIENCE,	// ��ѧ
	INDEX,		// ָ��
	END,		// ��β
	ERROR		// ����
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
				{ // ����
					state = INTEGER;
					bIsNumeric = true;
				}
				else if (s[i] == '+' || s[i] == '-')
				{ // ��ʼ
					state = START;
					bIsNumeric = false;
				}
				else if (s[i] == '.')
				{ // ������
					state = FLOAT;
					bIsNumeric = false;
				}
				else if (s[i] == ' ' || s[i] == '\t' || s[i] == '\n')
				{ // ���ǳ�ʼ
					state = INIT;
					bIsNumeric = false;
				}
				else
				{ // ����
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