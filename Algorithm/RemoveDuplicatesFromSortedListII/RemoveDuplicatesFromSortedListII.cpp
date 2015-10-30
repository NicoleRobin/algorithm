#include <iostream>

using namespace std;

/*
题意：
给定一个有序链表，删除有重复值的节点
思路：
遍历链表，保存该指向当前节点前一个节点的指针，如果当前节点和后续节点值相等，则删除当前节点及后续节点。
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
		ListNode *pPre = head;
		ListNode *pCur = head;
		while (pCur)
		{
			while (pCur->next != NULL && pCur->val == pCur->next->val)
			{
				pCur = pCur->next;
			}
			if (pPre == head)
			{
				pPre = pCur->next;
			}
			else
			{
				pPre->next = pCur->next;
			}
		}
		return head;
	}
};

int main(void)
{
	return 0;
}