#include <iostream>
#include <cstring>
#include <cmath>
#include <string>

using namespace std;

void _26(char *a, string str, int n)
{
	if (n == 0)
	{ // print
		if (str.empty())
		{
			cout << "Empty" << endl;
		}
		else
		{
			cout << str << endl;
		}
	}
	else
	{
		_26(a + 1, str, n - 1);
		str += a[0];
		_26(a + 1, str, n - 1);
	}
}

void _print_1(char *a, int i)
{
	int iIdx = 0;
	while (i > 0)
	{
		if (i % 2 == 1)
		{
			cout << a[iIdx];
		}
		i = i / 2;
		iIdx++;
	}
	cout << endl;
}

void _26_1(char *a, int n)
{
	int iLen = strlen(a);
	for (int i = 0; i < pow(2, iLen); ++i)
	{
		_print_1(a, i);
	}
}

int main(void)
{
	char a[] = "abc";
	string str;
	// _26(a, str, 3);
	_26_1(a, 3);
	return 0;
}
