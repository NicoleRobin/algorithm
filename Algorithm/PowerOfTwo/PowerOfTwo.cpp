#include <iostream>
#include <bitset>

using namespace std;

/*
题意：
给定一个数，判断是不是2的幂
思路：
首先需要注意的是，给定的数是整数，也就是可能是负数，而负数是不可能是2的幂的。
0也不可能是2的幂。
将给的整数转换为2进制形式，只有一个1的就是2的幂。
*/
class Solution {
public:
	bool isPowerOfTwo(int n) {
		bool bRet = false;
		if (n > 0)
		{
			bitset<32> bitN(n);
			if (bitN.count() == 1)
			{
				bRet = true;
			}
		}
		return bRet;
	}
};

int main(void)
{
	

	return 0;
}