#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	bool IsPopOrder(vector<int> pushV, vector<int> popV) {
		vector<int> vecTemp;
		
		int iPushLen = pushV.size();
		int iPopLen = popV.size();
		int iPushIdx = 0;
		int iPopIdx = 0;
		for (int iPopIdx = 0; iPopIdx < iPopLen; ++iPopIdx) {
			int iVal = popV[iPopIdx];

			if (vecTemp.empty() || vecTemp.back() != iVal) {
				for (int j = iPushIdx; j < iPushLen; ++j) {
					vecTemp.push_back(pushV[j]);
					if (pushV[j] == iVal) {
						iPushIdx = j + 1;
						break;
					}
				}
			}

			if (vecTemp.back() == iVal) {
				vecTemp.pop_back();
			}
			else {
				break;
			}
		}

		if (iPopIdx < iPopLen) {
			return false;
		}

		return true;
	}
};

int main(void)
{
	vector<int> vecPush;
	vecPush.push_back(1);
	vecPush.push_back(2);
	vecPush.push_back(3);
	vecPush.push_back(4);
	vecPush.push_back(5);
	vector<int> vecPop;
	vecPop.push_back(4);
	vecPop.push_back(5);
	vecPop.push_back(3);
	vecPop.push_back(2);
	vecPop.push_back(1);

	Solution s;
	cout << s.IsPopOrder(vecPush, vecPop) << endl;

	return 0;
}
