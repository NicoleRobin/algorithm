/*
 * g++ -g -o FindGreatestSumOfSubArray ./FindGreatestSumOfSubArray.cpp
 */
#include <iostream>
#include <vector>
using namespace std;

// 连续子数组的最大和
class Solution {
public:
	int FindGreatestSumOfSubArray(vector<int> array) {
		int iMaxSum = 0x80000000;
		cout << iMaxSum << endl;
		int iSum = 0;
		for (int i = 0; i < array.size(); ++i) {
			iSum += array[i];
			if (iMaxSum < iSum) {
				iMaxSum = iSum;
			}

			if (iSum <= 0) {
				iSum = 0;
			}
		}

		return iMaxSum;
	}
};

int main(void)
{
	vector<int> array;
	array.push_back(-1);
	array.push_back(-2);
	array.push_back(-3);
	array.push_back(-10);
	array.push_back(-4);
	array.push_back(-7);
	array.push_back(-2);
	array.push_back(-5);
	Solution s;
	cout << s.FindGreatestSumOfSubArray(array) << endl;

	return 0;
}
