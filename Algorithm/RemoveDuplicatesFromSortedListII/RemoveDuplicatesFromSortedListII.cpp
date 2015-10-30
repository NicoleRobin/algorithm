#include <iostream>

using namespace std;

/*
题意：
给定一个有序链表，删除有重复值的节点
思路：
遍历链表，保存该指向当前节点（pCur）和前一个节点（pPre）的指针，如果当前节点和后继节点相等，
则让当前节点指针继续向下走，直到与后继节点不相等，或者没有后继节点。
添加一个bool变量来表示当前节点是否和后继节点相等，
如果相等，则pPre->next = pCur->next; pCur = pCur->next;
如果不相等，则pPre = pPre->next; pCur = pCur->next;
*/
// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	ListNode* deleteDuplicates(ListNode* head) {
		if (head == NULL || head->next == NULL)
		{
			return head;
		}

		ListNode *pPre = new ListNode(0);
		pPre->next = head;
		ListNode *pCur = head;
		head = pPre;
		bool bDuplicated;
		while (pCur)
		{
			bDuplicated = false;
			while (pCur->next != NULL && pCur->val == pCur->next->val)
			{
				bDuplicated = true;
				pCur = pCur->next;
			}

			if (bDuplicated)
			{
				pPre->next = pCur->next;
				pCur = pCur->next;
			}
			else
			{
				pCur = pCur->next;
				pPre = pPre->next;
			}
		}
		return head->next;
	}
};

int main(void)
{
	ListNode *head = new ListNode(1);
	ListNode *p = head;
	p->next = new ListNode(1);
	p = p->next;
	p->next = new ListNode(1);
	p = p->next;
	p->next = new ListNode(2);
	p = p->next;
	p->next = new ListNode(3);

	Solution s;

	ListNode *pRet = s.deleteDuplicates(head);

	while (pRet)
	{
		if (pRet->next != NULL)
		{
			cout << pRet->val << "->";
		}
		else
		{
			cout << pRet->val;
		}
		pRet = pRet->next;
	}
	cout << endl;

	return 0;
}