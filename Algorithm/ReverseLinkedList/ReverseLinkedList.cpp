#include <iostream>
using namespace std;

/*
���⣺
����������
˼·��
��������������õݹ�ķ�����Ҳ�����õ����ķ�����
1���ݹ飬�ݹ�Ļ�ȡ��ÿһ���ڵ㣬Ȼ��β�巨�ؽ�����
2�������������Ļ�ȡ��ÿһ���ڵ㣬Ȼ��ͷ�巨�ؽ�����
*/
// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

// �ݹ�
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

// ����
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