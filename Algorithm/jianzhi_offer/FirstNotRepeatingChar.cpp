/*
 * g++ -g -o FirstNotRepeatingChar ./FirstNotRepeatingChar.cpp
 */
#include <iostream>
#include <string>
#include <iterator>
using namespace std;

class Solution {
public:
	int FirstNotRepeatingChar(string str) {
		int count[256] = {0};
		for (size_t i = 0; i < str.length(); ++i) {
			count[str[i]]++;
		}
		std::copy(count, count + 256, ostream_iterator<int>(cout, ", "));
		cout << endl;

		size_t index = -1;
		for (size_t i = 0; i < str.length(); ++i) {
			if (count[str[i]] == 1) {
				index = i;
				break;
			}
		}
		return index;
	}
};

int main(void)
{
	string str = "google";
	Solution s;
	cout << s.FirstNotRepeatingChar(str) << endl;
	return 0;
}
