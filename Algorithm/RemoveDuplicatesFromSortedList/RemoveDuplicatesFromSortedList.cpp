#include <iostream>
using namespace std;

/*
题意：
给定一个有序的单链表，删除其中的重复数据
思路：
顺序遍历单链表，遇到下一个节点与此节点相等的，就删除下一个节点，注意边界情况
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