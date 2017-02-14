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

void CountingSort(int a[], int len)
{
	int *temp = new int[10];
	for (int i = 0; i < 10; ++i)
	{
		temp[i] = 0;
	}

	// 统计每个数字的个数
	for (int i = 0; i < len; ++i)
	{
		temp[a[i]]++;
	}

	// 累加数字个数a[i] += a[i - 1]
	for (int i = 1; i < 10; ++i)
	{
		temp[i] += temp[i - 1];
	}

	int val = 0, idx = 0;
	int *temp1 = new int[10];
	for (int i = len - 1; i >= 0; i --)
	{
		val = a[i];
		idx = temp[val] - 1;
		temp1[idx] = val;
		temp[val]--;
	}

	for (int i = 0; i < len; ++i)
	{
		a[i] = temp1[i];
	}

	delete[] temp;
	delete[] temp1;
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 0, 8, 1, 4};
	Print(a, 10);

	// counting sort
	CountingSort(a, 10);
	Print(a, 10);

	return 0;
}
