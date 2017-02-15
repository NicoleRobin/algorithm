#include <iostream>
using namespace std;

/*
题意：
删除一个单链表中指定value的所有元素
思路：
遍历单链表，保存一个指向前一个元素的指针，如果当前节点value等于要删除的value，则删除该节点
*/
// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	ListNode* removeElements(ListNode* head, int val) {
		while (head)
		{
			if (head->val == val)
			{
				head = head->next;
			}
			else
			{
				break;
			}
		}
		
		if (head == NULL || (head->next == NULL && head->val == val))
		{
			return NULL;
		}

		ListNode *pPre = head;
		ListNode *p = head->next;
		
		while (p)
		{
			if (p->val == val)
			{
				pPre->next = p->next;
			}
			else
			{
				pPre = pPre->next;
			}
			p = p->next;
		}
		return head;
	}
};

int main(void)
{


	return 0;
}