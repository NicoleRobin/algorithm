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

// 调整某一个元素
void AdjustHeap(int a[], int i, int len)
{
	int nChild = 2 * i + 1;
	int nTemp = a[i];
	while (nChild < len)
	{
		// 找到子节点中较大的一个
		if (nChild + 1 < len && a[nChild] < a[nChild + 1])
		{
			nChild++;
		}
		if (nTemp < a[nChild])
		{
			a[i] = a[nChild];
			a[nChild] = nTemp;
			i = nChild;
			nChild = 2 * nChild + 1;
		}
		else
		{
			break;
		}
	}
}

// 初始化堆
void InitHeap(int a[], int len)
{
	for (int i = (len - 1) / 2; i >= 0; i--)
	{
		AdjustHeap(a, i, len);
	}
}

// 堆排序
void HeapSort(int a[], int len)
{
	// 初始堆
	InitHeap(a, len);
	// 排序
	int nTemp = 0;
	for (int i = len - 1; i > 0; i--)
	{
		// 交换第一个元素和最后一个元素
		nTemp = a[i]; a[i] = a[0]; a[0] = nTemp;

		// 调整堆
		AdjustHeap(a, 0, i);
	}
}

int main(void)
{
	int a[] = { 9, 7, 2, 6, 5, 3, 10, 0, 1, 4 };
	Print(a, 10);

	// select sort
	HeapSort(a, 10);
	Print(a, 10);

	return 0;
}