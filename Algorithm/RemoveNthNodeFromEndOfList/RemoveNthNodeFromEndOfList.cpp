#include <iostream>

using namespace std;

/*
���⣺
ɾ�������������е�����n���ڵ�
˼·��
��򵥵��뷨�����ȱ���һ����㵥����ĸ�����Ȼ�����Ҫɾ�����ǵڼ����ڵ㼴�ɡ�
���о��ǽ�����������ɾ����n��Ԫ�أ�Ȼ�������������

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

		// ����ڵ����
		int nLen = 0;
		ListNode *p = head;
		while (p)
		{
			nLen++;
			p = p->next;
		}

		// �����Ҫɾ�������ڼ����ڵ�
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