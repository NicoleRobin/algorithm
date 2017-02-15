#include <iostream>
using namespace std;

/*
题意：
判断给定的两个单链表是否有交叉点
思路：
如果两个单链表有交叉点，那么从交叉点之后的节点都一致，所以首先求出两个链表的长度。
假设长度差为n，那么从较长的第n个节点开始与较短的第一个节点对比，直到结尾。
*/

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
		int nLenA = 0, nLenB = 0;
		ListNode *p = headA;
		while (p)
		{
			nLenA++;
			p = p->next;
		}
		p = headB;
		while (p)
		{
			nLenB++;
			p = p->next;
		}

		ListNode *pRet = NULL;
		ListNode *q = headB;
		p = headA;
		if (nLenA > nLenB)
		{
			int nRet = nLenA - nLenB;
			while (nRet--)
			{
				p = p->next;
			}
		}
		else if (nLenA < nLenB)
		{
			int nRet = nLenB - nLenA;
			while (nRet--)
			{
				q = q->next;
			}
		}

		while (p && q)
		{
			if (p == q)
			{
				pRet = p;
				break;
			}
			p = p->next;
			q = q->next;
		}
		return pRet;
	}
};

int main(void)
{
	ListNode *headA = NULL;
	ListNode *headB = new ListNode(1);
	ListNode *p = headB;
	int n = 3;
	while (n < 22)
	{
		p->next = new ListNode(n);
		p = p->next;
		n += 2;
	}
	headA = p;
	p = headB;
	while (p)
	{
		if (p->next)
		{
			cout << p->val << "->";
		}
		else
		{
			cout << p->val << endl;
		}
		p = p->next;
	}
	Solution s;
	p = s.getIntersectionNode(headA, headB);
	if (p != NULL)
	{
		cout << p->val << endl;
	}
	else
	{
		cout << "NULL" << endl;
	}

	return 0;
}