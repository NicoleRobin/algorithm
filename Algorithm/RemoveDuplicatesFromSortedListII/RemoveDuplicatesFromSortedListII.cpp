#include <iostream>

using namespace std;

/*
���⣺
����һ����������ɾ�����ظ�ֵ�Ľڵ�
˼·��
�������������ָ��ǰ�ڵ�ǰһ���ڵ��ָ�룬�����ǰ�ڵ�ͺ����ڵ�ֵ��ȣ���ɾ����ǰ�ڵ㼰�����ڵ㡣
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