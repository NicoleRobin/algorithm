#include <iostream>
using namespace std;

/*
���⣺
ɾ�������������е�һ���ڵ㡣
˼·��
�������������ϸ��ڵ��ָ�룬����ɾ����ɾ��ָ���ڵ㣬ע��߽�������������
�޸ģ�
���������ˣ�����Ӧ���Ǹ���һ���ڵ㣬ɾ���ýڵ㣬��ԭ�ȵ�����Ǹ���һ������ɾ������ָ��value�Ľڵ㡣
����Ҳ�ͽ�����Ϊʲô��Ŀ���У�except the tail��given only access to that node�ˡ�
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