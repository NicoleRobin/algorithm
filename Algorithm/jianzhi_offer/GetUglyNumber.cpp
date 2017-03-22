#include <iostream>
#include <vector>
using namespace std;
// 把只包含因子2、3和5的数称作丑数（Ugly Number），1是第一个丑数
// 求第n个丑数
// 效率太慢了
class Solution1 {
public:
	bool IsUgly(int n) {
		while (n % 2 == 0) {
			n /= 2;
		}
		while (n % 3 == 0) {
			n /= 3;
		}
		while (n % 5 == 0) {
			n /= 5;
		}

		if (n == 1) {
			return true;
		}
		return false;
	}
	int GetUglyNumber_Solution(int index) {
		int i = 0;
		while (index > 0) {
			i++;
			if (IsUgly(i)) {
				index--;
			}
		}

		return i;
	}
};

/*
* 思路：计算出index个丑数，所有丑数都是由已有的丑数 *2 *3 *5得到的
*/
class Solution2 {
public:
	int GetUglyNumber_Solution(int index) {
		vector<int> vecUgly;
		vecUgly.push_back(1);
		int iUgly2, iUgly3, iUgly5, iMaxUgly, iMinUgly;
		while (vecUgly.size() <= index) {
			iMaxUgly = vecUgly[vecUgly.size() - 1];

			// 计算iUgly2
			for (int i = 0; i < vecUgly.size(); ++i) {
				if (vecUgly[i] * 2 > iMaxUgly) {
					iUgly2 = vecUgly[i] * 2;
					break;
				}
			}

			// 计算iUgly3
			for (int i = 0; i < vecUgly.size(); ++i) {
				if (vecUgly[i] * 3 > iMaxUgly) {
					iUgly3 = vecUgly[i] * 3;
					break;
				}
			}

			// 计算iUgly5
			for (int i = 0; i < vecUgly.size(); ++i) {
				if (vecUgly[i] * 5 > iMaxUgly) {
					iUgly5 = vecUgly[i] * 5;
					break;
				}
			}

			iMinUgly = iUgly2 <= iUgly3 ? iUgly2 : iUgly3;
			iMinUgly = iMinUgly <= iUgly5 ? iMinUgly : iUgly5;
			vecUgly.push_back(iMinUgly);
		}

		return vecUgly[index - 1];
	}
};

int main(void)
{
	Solution2 s;
	cout << s.GetUglyNumber_Solution(1500) << endl;
	return 0;
}
