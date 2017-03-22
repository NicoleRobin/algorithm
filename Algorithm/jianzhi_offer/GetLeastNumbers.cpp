/*
 * g++ -g -o GetLeastNumbers ./GetLeastNumbers.cpp
 */
#include <iostream>
#include <vector>
#include <iterator>
using namespace std;
/*
* 顺序遍历数组，把每一个数字尝试插入临时的容器，临时的容器必须保证，如果size小于k，直接插入，如果等于k，保证插入新数字之后个数不变。
*/
class Solution {
public:
	vector<int> GetLeastNumbers_Solution(vector<int> input, int k) {
		if (input.size() < k) {
			return vector<int>();
		}

		if (input.size() == k) {
			return input;
		}

		vector<int> vecTemp;
		for (int i = 0; i < input.size(); ++i) {
			if (vecTemp.size() < k) {
				vecTemp.push_back(input[i]);
				// 使用插入排序
				int j = vecTemp.size() - 1;
				while (j > 0) {
					if (vecTemp[j] < vecTemp[j - 1]) {
						int iTemp = vecTemp[j];
						vecTemp[j] = vecTemp[j - 1];
						vecTemp[j - 1] = iTemp;
					}
					j--;
				}
			}
			else {
				if (input[i] >= vecTemp.back()) {
					continue;
				}
				else {
					vecTemp[vecTemp.size() - 1] = input[i];
					// 使用插入排序
					int j = vecTemp.size() - 1;
					while (j > 0) {
						if (vecTemp[j] < vecTemp[j - 1]) {
							int iTemp = vecTemp[j];
							vecTemp[j] = vecTemp[j - 1];
							vecTemp[j - 1] = iTemp;
						}
						j--;
					}
				}
			}
			std::copy(vecTemp.begin(), vecTemp.end(), ostream_iterator<int>(cout, ","));
			cout << endl;
		}

		return vecTemp;
	}
};

int main(void)
{
	Solution s;
	vector<int> vecTemp;
	vecTemp.push_back(4);
	vecTemp.push_back(5);
	vecTemp.push_back(7);
	vecTemp.push_back(6);
	vecTemp.push_back(2);
	vecTemp.push_back(1);
	vecTemp.push_back(3);
	vecTemp.push_back(8);
	vector<int> vecRet = s.GetLeastNumbers_Solution(vecTemp, 5);

	std::copy(vecRet.begin(), vecRet.end(), ostream_iterator<int>(cout, ","));
	cout << endl;

	return 0;
}
