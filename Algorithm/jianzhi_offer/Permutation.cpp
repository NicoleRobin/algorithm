/*
 * g++ -g --std=c++11 -o Permutation ./Permutation.cpp
 */
#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Solution {
public:
	vector<string> Permutation(string str) {
		if (str.length() == 0)
		{
			return vector<string>();
		}

		string strPath;
		vector<string> vecRet;
		Permutation(str, strPath, vecRet);
		return vecRet;
	}

	void Permutation(string str, string strPath, vector<string> &vecRet)
	{
		if (str.length() == 0)
		{
			vecRet.push_back(strPath);
			return;
		}

		for (int i = 0; i < str.length(); ++i)
		{
			if (i > 0 && str[i] == str[0])
			{
				continue;
			}
			// 牛客网oj的锅，同样的输出，顺序不一样不能通过，因此这么写
			/*
			char ch = str[0];
			str[0] = str[i];
			str[i] = ch;
			*/
			cout << str << endl;
			char ch = str[i];
			str.erase(i, 1);
			str.insert(str.begin(), ch);
			cout << str << endl;

			strPath.push_back(str[0]);
			Permutation(string(str.begin() + 1, str.end()), strPath, vecRet);
			strPath.pop_back();
			/*
			ch = str[0];
			str[0] = str[i];
			str[i] = ch;
			*/
			ch = str[0];
			str.erase(0, 1);
			str.insert(str.begin() + i, ch);
			cout << str << endl;
		}
	}
};

int main(void)
{
	Solution s;
	string str="aa";
	vector<string> vecRet = s.Permutation(str);
	for (int i = 0; i < vecRet.size(); ++i)
	{
		cout << vecRet[i] << endl;
	}

	return 0;
}
