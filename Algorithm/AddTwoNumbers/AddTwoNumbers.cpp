#include <iostream>
#include <vector>
#include <algorithm>
#include <fstream>
#include <iterator>

using namespace std;

/*
���⣺
�������������ʾ�ķǸ������������ǵĺͣ�������ķ�ʽ���ء�
�Ǹ������е�ÿһλ�����ֵ�˳���Ƿ���洢�������еġ�
˼·��
ѭ����������ͬʱ���Ӧλ�ϵ��������ֵĺͣ��������β�巨����һ���µ������С�
��Ҫע����Ƕ����һ����λ�Ĵ�����Ҫ�����ˡ�
*/

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
		ListNode *pList = new ListNode(0);
		ListNode *qList = pList;
		int nTemp = 0;
		while (l1 != NULL || l2 != NULL)
		{
			if (l1 != NULL)
			{
				nTemp += l1->val;
				l1 = l1->next;
			}
			if (l2 != NULL)
			{
				nTemp += l2->val;
				l2 = l2->next;
			}
			pList->next = new ListNode(nTemp % 10);
			nTemp = nTemp / 10;
			pList = pList->next;
		}
		if (nTemp != 0)
		{
			pList->next = new ListNode(nTemp);
		}

		pList = qList->next;
		delete qList;
		return pList;
	}
};

int main()
{
	Solution solution;
	ListNode *p = NULL, *q = NULL;
	p = new ListNode(2);
	p->next = new ListNode(4);
	p->next->next = new ListNode(3);
	q = new ListNode(5);
	q->next = new ListNode(6);
	q->next->next = new ListNode(4);

	ListNode *pRet = solution.addTwoNumbers(p, q);

	while (pRet != NULL)
	{
		if (pRet->next == NULL)
		{
			cout << pRet->val << endl;
		}
		else
		{
			cout << pRet->val << "->";
		}
		pRet = pRet->next;
	}

	return 0;
}