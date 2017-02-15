#include <iostream>
using namespace std;

/*
���⣺
����һ������������Ƿ��ǻ��ġ�
˼·��
�ȱ���һ����������ڵ���������벿�����ж�ǰ�벿�����벿���Ƿ���ȣ���󽫺�벿�����������
*/

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

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

	bool isPalindrome(ListNode* head) {
		bool bRet = true;
		int nLen = 0;
		ListNode *p = head;
		while (p)
		{
			nLen++;
			p = p->next;
		}
		
		int nHalf = 0;
		if (nLen % 2 == 0)
		{
			nHalf = nLen / 2;
		}
		else
		{
			nHalf = nLen / 2 + 1;
		}
		ListNode *pFirstHalf = head, *pSecondHalf = NULL;
		p = head;
		while (p && nHalf-- > 0)
		{
			p = p->next;
		}
		pSecondHalf = p;
		ListNode *pOrigin = reverseList(pSecondHalf);
		pSecondHalf = pOrigin;
		while (pFirstHalf && pSecondHalf)
		{
			if (pFirstHalf->val != pSecondHalf->val)
			{
				bRet = false;
				break;
			}
			pFirstHalf = pFirstHalf->next;
			pSecondHalf = pSecondHalf->next;
		}
		reverseList(pOrigin);
		return bRet;
	}
};

int main(void)
{


	return 0;
}