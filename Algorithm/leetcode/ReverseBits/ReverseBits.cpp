#include <iostream>
#include <bitset>
using namespace std;

/*
���⣺
��һ���������޷�������ת��Ϊ���������ʽ��Ȼ��ת�����ƣ�����������ת��Ϊ�޷����������ء�
˼·��
����C++��bitset��
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