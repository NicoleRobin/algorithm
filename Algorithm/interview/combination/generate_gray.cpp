#include <iostream>
#include <iterator>
#include <vector>
#include <string>

using namespace std;

vector<string> GenerateGrayCode(int n)
{
	vector<string> grayCode;
	if (n == 1)
	{
		grayCode.push_back("0");
		grayCode.push_back("1");
		return grayCode;
	}

	grayCode = GenerateGrayCode(n - 1);
	vector<string> newGrayCode;
	for (int i = 0; i < grayCode.size(); ++i)
	{
		newGrayCode.push_back(grayCode[i] + "0");
		newGrayCode.push_back(grayCode[i] + "1");
	}
	
	return newGrayCode;
}



int main(void)
{

	vector<string> grayCode = GenerateGrayCode(4);
	copy(grayCode.begin(), grayCode.end(), ostream_iterator<string>(cout, "\n"));
	cout << endl;

	return 0;
}
