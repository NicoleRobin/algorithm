#include <iostream>
using namespace std;

/*
���⣺
��N��С��վ��һ�ţ�ÿ��С������һ��Ȩֵ����ҪһЩ�ǰ�����Ĺ���ָ���ЩС��������ʱ������Ҫ��
1��ÿ��С��������һ���ǡ�
2��ӵ�нϸ�Ȩֵ��С���������ھӻ�ø�����ǡ�
��������Ҫ�ö�����
˼·��

*/
class Solution {
public:
	int candy(vector<int>& ratings) {
		int nRet = ratings.size();
		for (int i = 0; i < ratings.size(); ++i)
		{
			if (i == 0)
			{
				if (ratings[i] > ratings[i + 1])
				nRet += 1;
			}
			else if (i == ratings.size() - 1)
			{
				if (ratings[i] > ratings[i - 1])
				{
					nRet += 1;
				}
			}
			else
			{
				if (ratings[i] > )
			}
		}
	}
};

int main(void)
{


	return 0;
}