#include <iostream>

using namespace std;

/*
题意：
合并两个有序的单链表
思路：
遍历两个单链表，比较大小决定合并的顺序
疑问：
没有说明两个有序链表是升序还是降序。不知道需不需要判断，还是默认升序处理？
*/
// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};
class Solution {
public:
	ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
		ListNode *pRet = NULL;
		if (l1 == NULL || l2 == NULL)
		{
			pRet = l1 == NULL ? l2 : l1;
		}
		else
		{
			if (l1->val < l2->val)
			{
				pRet = l1;
				l1 = l1->next;
			}
			else
			{
				pRet = l2;
				l2 = l2->next;
			}
			ListNode *p = pRet;
			while (l1 && l2)
			{
				if (l1->val < l2->val)
				{
					p->next = l1;
					p = p->next;
					l1 = l1->next;
				}
				else
				{
					p->next = l2;
					p = p->next;
					l2 = l2->next;
				}
			}
			p->next = l1 == NULL ? l2 : l1;
		}
		return pRet;
	}
};

int main(void)
{
	ListNode *l1 = new ListNode(1);
	ListNode *l2 = new ListNode(2);

	Solution s;
	ListNode *p = s.mergeTwoLists(l1, l2);
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

	return 0;
}