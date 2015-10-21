#include <iostream>
using namespace std;

/*
题意：
有N个小孩站成一排，每个小孩分配一个权值，你要一些糖按下面的规则分给这些小孩。分配时有以下要求：
1、每个小孩至少有一个糖。
2、拥有较高权值的小孩比他的邻居获得更多的糖。
求你最少要拿多少糖
思路：

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