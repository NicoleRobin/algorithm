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
		int temp;
		int k;
		temp = a[i];
		for (int j = i; j < len - 1; j++)
		{
			if (a[j] < temp)
			{
				temp = a[j];
				k = j;
			}
			if (k != i)
			{ // swap
				swap(&a[i], &temp);
				swap(&temp, &a[k]);
			}
		}
		
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