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

void MergeArray(int a[], int aLen, int b[], int bLen, int temp[])
{
	int i = 0, j = 0, k = 0;
	while (i < aLen && j < bLen)
	{
		if (a[i] < b[j])
		{
			temp[k++] = a[i++];
		}
		else
		{
			temp[k++] = b[j++];
		}
	}

	while (i < aLen)
	{
		temp[k++] = a[i++];
	}

	while (j < bLen)
	{
		temp[k++] = b[j++];
	}
}

void MergeSort(int a[], int n)
{
	if (n < 2)
	{
		return;
	}

	int m = n / 2;
	MergeSort(a, m);
	MergeSort(a + m, n - m);
	int *temp = new int[n];
	MergeArray(a, m, a + m, n - m, temp);
	for (int i = 0; i < n; ++i)
	{
		a[i] = temp[i];
	}
	delete[] temp;
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 10, 0, 1, 4 };
	Print(a, 10);

	MergeSort(a, 10);
	Print(a, 10);

	return 0;
}
