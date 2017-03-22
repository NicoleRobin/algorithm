/*
 * g++ -g -o VerifySequenceOfBST ./VerifySequenceOfBST.cpp
 */
#include <iostream>
#include <vector>
#include <iterator>
using namespace std;

class Solution {
public:
	bool VerifySequenceOfBST(vector<int> sequence) {
		if (sequence.size() == 0 || sequence.size() == 1) {
			return true;
		}
		std::copy(sequence.begin(), sequence.end(), ostream_iterator<int>(cout, ", "));
		cout << endl;

		int iLen = sequence.size();
		int iRoot = sequence[iLen - 1];
		vector<int> vecLeft;
		vector<int> vecRight;
		int i = 0;
		for (; i < iLen - 1; ++i) {
			if (sequence[i] > iRoot) {
				break;
			}
			vecLeft.push_back(sequence[i]);
		}
		for (; i < iLen - 1; ++i) {
			if (sequence[i] < iRoot) {
				return false;
			}
			vecRight.push_back(sequence[i]);
		}

		return VerifySequenceOfBST(vecLeft) && VerifySequenceOfBST(vecRight);
	}
};
int main(void)
{
	vector<int> vecBst;
	vecBst.push_back(1);
	vecBst.push_back(4);
	vecBst.push_back(3);
	vecBst.push_back(6);
	vecBst.push_back(8);
	vecBst.push_back(7);
	vecBst.push_back(5);

	Solution s;
	cout << s.VerifySequenceOfBST(vecBst) << endl;

	return 0;
}
