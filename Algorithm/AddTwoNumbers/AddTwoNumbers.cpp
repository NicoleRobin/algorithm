#include <iostream>
#include <vector>
#include <algorithm>
#include <fstream>
#include <iterator>

using namespace std;

/*
题意：
给定两个链表表示的非负整数，求它们的和，以链表的方式返回。
非负整数中的每一位的数字的顺序是反向存储在链表中的。
思路：
循环两个链表，同时求对应位上的两个数字的和，将结果按尾插法插入一个新的链表中。
需要注意的是对最后一个进位的处理。不要丢掉了。
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