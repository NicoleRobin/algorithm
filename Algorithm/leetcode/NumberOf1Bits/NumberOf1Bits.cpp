#include <iostream>
using namespace std;

/*
题意：
给定一个整数，算出它对应的二进制中1的个数。
思路：
对给定的一个整数先对2求余，看是否为1，然后除2，如此循环直至该数为0.
*/
class Solution {
public:
	int hammingWeight(unsigned int n) {
		int nCount = 0;
		while (n > 0)
		{
			if (n % 2 == 1)
			{
				++nCount;
			}
			n /= 2;
		}
		return nCount;
	}
};

int main(void)
{
	unsigned int n;
	Solution s;
	while (true)
	{
		cin >> n;
		cout << s.hammingWeight(n) << endl;
	}

	return 0;
}