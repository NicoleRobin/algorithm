#include <iostream>
#include <bitset>

using namespace std;

/*
���⣺
����һ�������ж��ǲ���2����
˼·��
������Ҫע����ǣ�����������������Ҳ���ǿ����Ǹ������������ǲ�������2���ݵġ�
0Ҳ��������2���ݡ�
����������ת��Ϊ2������ʽ��ֻ��һ��1�ľ���2���ݡ�
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