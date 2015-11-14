#include "string.h"
#include <iostream>

using namespace std;
int main(void)
{
	ZY::string strTest = "hello world";

	cout << strTest << endl;
	cout << "length:" << strTest.length() << endl;
	cout << "strTest[2]:" << strTest[2] << endl;
	return 0;
}