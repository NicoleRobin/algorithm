#include <stdio.h>
#include <stdlib.h>
#include <time.h>

// bubble sort
void zjw_sort(int num[], int len)
{
	int temp;
	for (int i = 0; i < len; i++)
	{
		for (int j = 0; j < len - i; j++)
		{
			if (num[j] > num[j + 1])
			{ // swap
				temp = num[j];
				num[j] = num[j + 1];
				num[j + 1] = temp;
			}
		}
	}
}

void yang_sort(int num[], int len){
	for (int i = 0; i < len; ++i){
		for (int j = len - 1; j > i; --j){
			if (num[j] < num[j - 1]){
				// swap(num[j], num[j - 1]);
			}
		}
	}
}
int main()
{
	int a[100];
	srand((unsigned)time(NULL));
	for (int i = 0; i < 100; i++)
	{
		a[i] = rand() % 100;
		printf("%d ", a[i]);
	}
	zjw_sort(a, 99);

	printf("\n");
	for (int i = 0; i < 100; i++)
	{
		printf("%d ", a[i]);
	}

	return 0;
}