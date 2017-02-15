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

void RadixSort(int a[], int len)
{
	int iMax = 0;
	for (int i = 0; i < len; ++i)
	{
		if (abs(a[i]) > iMax)
		{
			iMax = abs(a[i]);
		}
	}
	printf("Max[%d]\n", iMax);

	int iRadix = 0;
	while (iMax > 0)
	{
		iMax /= 10;
		iRadix++;
	}
	printf("Radix[%d]\n", iRadix);

	for (int i = 0; i < iRadix; ++i)
	{
		int iDiv = 1;
		for (int j = 0; j < i; ++j)
		{
			iDiv *= 10;
		}
		// 使用计数排序按每一位进行排序
		int counting[10] = {0};
		for (int j = 0; j < len; ++j)
		{
			counting[(a[j] / iDiv) % 10]++;
		}

		for (int j = 1; j < 10; ++j)
		{
			counting[j] += counting[j - 1];
		}

		int val = 0, idx = 0;
		int *temp = new int[len];
		for (int j = len - 1;j >= 0; --j)
		{
			val = (a[j] / iDiv) % 10;
			idx = counting[val] - 1;
			temp[idx] = a[j];
			counting[val]--;
		}

		for (int j = 0; j < len; ++j)
		{
			a[j] = temp[j];
		}

		delete[] temp;
		Print(a, len);
	}
}

int main(void)
{
	int a[] = { 326, 453, 608, 835, 751, 435, 704, 690, 88, 79 };
	Print(a, 10);

	RadixSort(a, 10);
	Print(a, 10);

	return 0;
}
