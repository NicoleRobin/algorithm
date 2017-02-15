#include <iostream>
using namespace std;

/*
题意：
删除给定单链表中的一个节点。
思路：
遍历单链表保存上个节点的指针，用于删除，删除指定节点，注意边界条件（空链表）
修改：
理解错题意了，题意应该是给定一个节点，删除该节点，我原先的理解是给定一个链表，删除其中指定value的节点。
这样也就解释了为什么题目中有：except the tail和given only access to that node了。
*/
// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

class Solution {
public:
	void deleteNode(ListNode* node) {
		if (node == NULL)
		{
			return;
		}
		node->val = node->next->val;
		node->next = node->next->next;
	}
};

int main(void)
{


	return 0;
}