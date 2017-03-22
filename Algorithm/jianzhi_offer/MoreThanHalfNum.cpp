/*
* g++ -g -o MoreThanHalfNum ./MoreThanHalfNum.cpp
*/
#include <iostream>
#include <vector>
using namespace std;

/*
* 根据数组的特点，得出解法：遍历数组，记录数字和次数，如果当前数字和之前记录的数字相同就对次数加一，否则将次数减一，当减为0的时候
* 记录数字替换为当前数字，并重置次数为1
*/
class Solution {
public:
    int MoreThanHalfNum_Solution(vector<int> numbers) {
    	if (numbers.size() <= 0) {
            return 0;
        }
        
        int result = numbers[0];
        int count = 1;
        for (int i = 1; i < numbers.size(); ++i) {
            if (numbers[i] == result) {
                count++;
            }
            else {
                count--;
                if (count == 0) {
                    result = numbers[i];
                    count = 1;
                }
            }
        }
        
        // 检查result是否次数超过长度一半
	count = 0;
        for (int i = 0; i < numbers.size(); ++i) {
            if (numbers[i] == result) {
                count++;
            }
        }
        
        if (count <= numbers.size() / 2) {
            return 0;
        }
        
        return result;
    }
};

int main(void)
{
	vector<int> numbers;
	numbers.push_back(1);
	numbers.push_back(1);
	numbers.push_back(1);
	numbers.push_back(1);
	numbers.push_back(1);
	numbers.push_back(1);

	Solution s;
	cout << s.MoreThanHalfNum_Solution(numbers) << endl;
	
	return 0;
}
