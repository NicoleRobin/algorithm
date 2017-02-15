#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include <iterator>
#include <list>

using namespace std;

void Print(int a[], int len)
{
	for (int i = 0; i < len; i++)
	{
		printf("%d ", a[i]);
	}
	printf("\n");
}

void swap(int* x, int* y)
{
	int temp;
	temp = *x;
	*x = *y;
	*y = temp;
}

void BucketSort(int a[], int len)
{
	std::list<int> temp[10];
	for (int i = 0; i < len; ++i)
	{
		int idx = a[i] / 10;
		temp[idx].push_back(a[i]);
	}

	cout << "Before sort" << endl;
	for (int i = 0; i < 10; ++i)
	{
		cout << "Bucket " << i << "\t";
		std::copy(temp[i].begin(), temp[i].end(), ostream_iterator<int>(cout, " "));
		cout << endl;
	}
	cout << endl;

	int j = 0;
	for (int i = 0; i < 10; ++i)
	{
		temp[i].sort();
		while (!temp[i].empty())
		{
			a[j++] = temp[i].front();
			temp[i].pop_front();
		}
	}
}

int main(void)
{
	int a[] = { 92, 71, 27, 63, 5, 39, 10, 0, 16, 42};
	Print(a, 10);

	// Bucket Sort
	BucketSort(a, 10);
	Print(a, 10);

	return 0;
}
