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
	for (int i = 0; i < n; ++i)
	{
		if (abs(a[i]) > iMax)
		{
			iMax = abs(a[i]);
		}
	}

	int iRadix = 0;
	while (iMax > 0)
	{
		iMax /= 10;
		iRadix++;
	}

	for (int i = 0; i < iRadix; ++i)
	{
		
	}
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 10, 0, 1, 4 };
	Print(a, 10);

	RadixSort(a, 10);
	Print(a, 10);

	return 0;
}
