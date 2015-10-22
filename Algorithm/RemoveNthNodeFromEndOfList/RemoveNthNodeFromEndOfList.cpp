#include <iostream>

using namespace std;

/*
题意：
删除给定单链表中倒数第n个节点
思路：
最简单的想法就是先遍历一遍计算单链表的个数，然后算出要删除的是第几个节点即可。
还有就是将单链表逆序，删除第n个元素，然后再逆序回来。

*/
// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	ListNode* removeNthFromEnd(ListNode* head, int n) {
		if (head == NULL || n <= 0)
		{
			return head;
		}

		// 计算节点个数
		int nLen = 0;
		ListNode *p = head;
		while (p)
		{
			nLen++;
			p = p->next;
		}

		// 算出是要删除正数第几个节点
		int nIndex = nLen - n + 1;

		if (nIndex == 1)
		{
			head = head->next;
		}
		else
		{
			nIndex--;
			ListNode *pPre = head;
			while (--nIndex)
			{
				pPre = pPre->next;
			}
			pPre->next = pPre->next->next;
		}
		return head;
	}
};

int main(void)
{


	return 0;
}