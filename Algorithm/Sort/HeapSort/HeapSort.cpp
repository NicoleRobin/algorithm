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

// ����ĳһ��Ԫ��
void AdjustHeap(int a[], int i, int len)
{
	int nChild = 2 * i + 1;
	int nTemp = a[i];
	while (nChild < len)
	{
		// �ҵ��ӽڵ��нϴ��һ��
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

// ��ʼ����
void InitHeap(int a[], int len)
{
	for (int i = (len - 1) / 2; i >= 0; i--)
	{
		AdjustHeap(a, i, len);
	}
}

// ������
void HeapSort(int a[], int len)
{
	// ��ʼ��
	InitHeap(a, len);
	// ����
	int nTemp = 0;
	for (int i = len - 1; i > 0; i--)
	{
		// ������һ��Ԫ�غ����һ��Ԫ��
		nTemp = a[i]; a[i] = a[0]; a[0] = nTemp;

		// ������
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