#include <iostream>

using namespace std;

/*
���⣺
����һ����������ɾ�����ظ�ֵ�Ľڵ�
˼·��
�������������ָ��ǰ�ڵ㣨pCur����ǰһ���ڵ㣨pPre����ָ�룬�����ǰ�ڵ�ͺ�̽ڵ���ȣ�
���õ�ǰ�ڵ�ָ����������ߣ�ֱ�����̽ڵ㲻��ȣ�����û�к�̽ڵ㡣
���һ��bool��������ʾ��ǰ�ڵ��Ƿ�ͺ�̽ڵ���ȣ�
�����ȣ���pPre->next = pCur->next; pCur = pCur->next;
�������ȣ���pPre = pPre->next; pCur = pCur->next;
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