#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
	void FindNumsAppearOnce(vector<int> data, int* num1, int *num2) {
		if (data.size() < 2){
			return;
		}

		int iXOR = data[0];
		for (size_t i = 1; i < data.size(); ++i) {
			iXOR ^= data[i];
			cout << iXOR << endl;
		}

		bool bNum1Find = false;
		for (size_t i = 0; i < data.size(); ++i) {
			if (iXOR ^ data[i] == 0) {
				if (!bNum1Find) {
					*num1 = data[i];
				}
				else
				{
					*num2 = data[i];
					break;
				}
			}
			
		}
	}
};

int main(void)
{
	vector<int> vecTest;
	vecTest.push_back(1);
	vecTest.push_back(1);
	vecTest.push_back(2);
	vecTest.push_back(3);
	vecTest.push_back(4);
	vecTest.push_back(4);

	Solution s;
	int iNum1, iNum2;
	s.FindNumsAppearOnce(vecTest, &iNum1, &iNum2);
	cout << "Num1:" << iNum1 << ", Num2:" << iNum2 << endl;

	return 0;
}
