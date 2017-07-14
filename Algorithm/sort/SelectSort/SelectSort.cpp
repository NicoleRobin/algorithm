/*
* g++ -g -o SelectSort ./SelectSort.cpp
*/
#include <stdio.h>
#include <stdlib.h>

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

void SelectSort(int a[], int len)
{
	for (int i = 0; i < len; i++)
	{
		int iMin = a[i];
		int iMinIdx = i;
		for (int j = i; j < len; j++)
		{
			if (a[j] < iMin)
			{
				iMin = a[j];
				iMinIdx = j;
			}
		}
		if (iMinIdx != i)
		{ // swap
			swap(&a[i], &a[iMinIdx]);
		}
		Print(a, 10);	
	}
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 10, 0, 1, 4 };
	Print(a, 10);

	// select sort
	SelectSort(a, 10);
	Print(a, 10);

	return 0;
}
