#include <iostream>
using namespace std;

/*
���⣺
����һ������ĵ�����ɾ�����е��ظ�����
˼·��
˳�����������������һ���ڵ���˽ڵ���ȵģ���ɾ����һ���ڵ㣬ע��߽����
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
		ListNode* pRet = head;
		while (head && head->next)
		{
			if (head->val == head->next->val)
			{
				head->next = head->next->next;
			}
			else
			{
				head = head->next;
			}
		}
		return pRet;
	}
};

int main(void)
{
	int n;
	ListNode *pRoot = new ListNode(0);
	ListNode *p = pRoot;
	while (cin >> n)
	{
		p->next = new ListNode(n);
		p = p->next;
	}
	Solution s;
	s.deleteDuplicates(pRoot->next);
	p = pRoot->next;
	while (p)
	{
		if (p->next)
		{
			cout << p->val << "->"
		}
		else
		{
			cout << p->val << endl;
		}
	}

	return 0;
}