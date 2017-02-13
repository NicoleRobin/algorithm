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

void MergeArray(int a[], int aLen, int b[], int bLen)
{
	if (aLen == 0 bLen == 0)
	{
	}
}

void MergeSort(int a[], int len)
{
	
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 10, 0, 1, 4 };
	Print(a, 10);

	MergeSort(a, 10);
	Print(a, 10);

	return 0;
}
