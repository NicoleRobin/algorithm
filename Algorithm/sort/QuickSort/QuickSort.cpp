#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void QuickSort(int a[], int nStart, int nEnd)
{
	if (nStart < nEnd)
	{
		int x = a[nStart];
		int y = nStart;
		int i = nStart, j = nEnd;
		while (i < j)
		{
			while (i < j && a[j] >= x)
			{
				j--;
			}
			a[y] = a[j];
			y = j;

			while (i < j && a[i] <= x)
			{
				i++;
			}
			a[y] = a[i];
			y = i;
		}
		a[y] = x;

		QuickSort(a, nStart, y - 1);
		QuickSort(a, y + 1, nEnd);
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
	QuickSort(a, 0, 99);

	printf("\n");
	for (int i = 0; i < 100; i++)
	{
		printf("%d ", a[i]);
	}

	return 0;
}

void swap(int &i, int &j){
	int temp = i;
	i = j;
	j = temp;
}
int sort(int num[], int low, int high){
	int temp = num[low];
	while (low < high){
		while (num[low] < temp && low <= high) ++low;
		while (num[high] > temp && low <= high) --high;
		swap(num[low], num[high]);
	}
	return low;
}

void quickSort(int num[], int low, int high){
	if (low < high){
		int pos = sort(num, low, high);
		quickSort(num, low, pos - 1);
		quickSort(num, pos + 1, high);
	}
}


int main(){

	int a[100];
	srand((unsigned)time(NULL));
	for (int i = 0; i < 100; i++)
	{
		a[i] = rand() % 100;
		printf("%d ", a[i]);
	}
	quickSort(a, 0, 99);

	printf("\n");
	for (int i = 0; i < 100; i++)
	{
		printf("%d ", a[i]);
	}
	return 0;
}