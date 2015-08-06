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

void InsertSort(int a[], int len)
{
	for (int i = 1; i < len; i++)
	{
		if (a[i] < a[i - 1])
		{
			int j = i;
			while (j > 0)
			{
				if (a[j] < a[j - 1])
				{ // ½»»»
					swap(&a[j], &a[j - 1]);
				}
				else
				{
					break;
				}
				j--;
			}
		}
	}
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 10, 0, 1, 4};
	Print(a, 10);

	// insert sort
	InsertSort(a, 10);
	Print(a, 10);

	return 0;
}