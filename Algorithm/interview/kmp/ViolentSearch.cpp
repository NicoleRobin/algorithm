/*
* g++ -g -o ViolentSearch ./ViolentSearch.cpp
*/
#include <iostream>
#include <cstring>

using namespace std;

bool ViolentSearch(char *s, char *p)
{
	int iSLen = strlen(s);
	int iPLen = strlen(p);

	if (iSLen < iPLen)
	{
		return false;
	}

	bool bEqual = true;
	for (int i = 0; i < iSLen; ++i)
	{
		bEqual = true;
		for (int j = 0; j < iPLen; ++j)
		{
			if (i + j >= iSLen)
			{
				break;
			}

			if (s[i + j] != p[j])
			{
				bEqual = false;
				break;
			}
			
		}

		if (bEqual)
		{
			break;
		}
	}

	return bEqual;
}

bool ViolentSearch2(char *s, char *p)
{
	int iSLen = strlen(s);
	int iPLen = strlen(p);

	int i = 0;
	int j = 0;
	while (i < iSLen && j < iPLen)
	{
		if (s[i] == p[j])
		{
			i++;
			j++;
		}
		else
		{
			// i和j分别后退j步，同时i向前进1步
			i = i - j + 1;
			j = 0;
		}
	}

	if (j == iPLen)
	{
		return true;
	}
	else
	{
		return false;
	}
}

int main(void)
{
	char s[] = "Hello world!";
	char p[] = "lld";

	if (ViolentSearch(s, p))
	{
		cout << "Match" << endl;
	}
	else
	{
		cout << "Unmatch" << endl;
	}

	return 0;
}
