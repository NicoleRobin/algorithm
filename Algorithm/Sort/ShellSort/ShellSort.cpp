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

void ShellSort(int a[], int len)
{
	int interval = len / 2;
	while (interval > 0)
	{
		for (int i = 0; i < interval; i++)
		{
			for (int j = i + interval; j < len;j += interval)
			{
				int k = j;
				while (k > 0)
				{
					if (a[k] < a[k - interval])
					{
						swap(&a[k], &a[k - interval]);
					}
					else
					{
						break;
					}
					k = k - interval;
				}
			}
		}
		interval = interval / 2;
	}
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 10, 0, 1, 4 };
	Print(a, 10);

	// select sort
	ShellSort(a, 10);
	Print(a, 10);

	return 0;
}