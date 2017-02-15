#include <iostream>
using namespace std;

/*
题意：
单链表逆序
思路：
单链表逆序可以用递归的方法，也可以用迭代的方法。
1、递归，递归的获取到每一个节点，然后按尾插法重建链表。
2、迭代，遍历的获取到每一个节点，然后按头插法重建链表。
*/
// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// 递归
class Solution {
public:
	ListNode* reverseList(ListNode* head) {
		if (head == NULL)
		{
			return NULL;
		}
		if (head->next == NULL)
		{
			return head;
		}
		else
		{
			ListNode *pRet = NULL;
			pRet = reverseList(head->next);
			head->next = NULL;
			
			ListNode *pLast = pRet;
			while (pLast->next)
			{
				pLast = pLast->next;
			}
			pLast->next = head;
			return pRet;
		}
	}
};

// 迭代
class Solution {
public:
	ListNode* reverseList(ListNode* head) {
		if (head == NULL)
		{
			return NULL;
		}

		ListNode *pRet = head;
		head = head->next;
		pRet->next = NULL;
		ListNode *pNext = NULL;
		while (head)
		{
			pNext = head->next;
			head->next = pRet;
			pRet = head;
			head = pNext;
		}
		return pRet;
	}
};

int main(void)
{
	ListNode *p = new ListNode(1);
	p->next = new ListNode(2);
	p->next->next = new ListNode(3);

	Solution s;
	ListNode *pRet = s.reverseList(p);

	while (pRet)
	{
		if (pRet->next)
		{
			cout << pRet->val << "->";
		}
		else
		{
			cout << pRet->val << endl;
		}
		pRet = pRet->next;
	}

	return 0;
}