#include <iostream>
#include <bitset>
using namespace std;

/*
题意：
将一个给定的无符号整数转换为其二进制形式，然后翻转二进制，并将二进制转换为无符号整数返回。
思路：
利用C++的bitset。
*/
typedef unsigned int uint32_t;

class Solution {
public:
	uint32_t reverseBits(uint32_t n) {
		bitset<32> bitN(n);
		string strN = bitN.to_string();
		reverse(strN.begin(), strN.end());
		bitset<32> bitRet(strN);
		return (uint32_t)bitRet.to_ulong();
	}
};

int main(void)
{
	Solution s;
	uint32_t n;
	while (true)
	{
		cin >> n;
		cout << s.reverseBits(n) << endl;
	}

	return 0;
}