#include <iostream>
using namespace std;

/*
���⣺
ɾ��һ����������ָ��value������Ԫ��
˼·��
��������������һ��ָ��ǰһ��Ԫ�ص�ָ�룬�����ǰ�ڵ�value����Ҫɾ����value����ɾ���ýڵ�
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