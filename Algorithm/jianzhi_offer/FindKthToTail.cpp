/*
* g++ -g -o FindKthToTail ./FindKthToTail.cpp
*/
#include <iostream>
using namespace std;

struct ListNode {
	int val;
	struct ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};

class Solution {
public:
    ListNode* FindKthToTail(ListNode* pListHead, unsigned int k) {
    	if (pListHead == NULL) {
            return pListHead;
        }
        
        // 当第一个指针走到第k个节点时，第二个节点开始走
        // 当第一个指针走到结尾时，第二个指针指向倒数第k个节点
        int i = 0;
        ListNode *pFirst = pListHead;
        ListNode *pSecond = NULL;
        while (pFirst) {
            i++;
            if (i == k) {
                pSecond = pListHead;
            }
            else if (i > k) {
                pSecond = pSecond->next;
            }
            pFirst = pFirst->next;
        }
        
        return pSecond;
    }
};

int main(void)
{	
	ListNode *pHead = new ListNode(1);
	ListNode *pNext = pHead;
	for (int i = 2; i < 6; ++i)
	{
		ListNode *pNode = new ListNode(i);
		pNext->next = pNode;
		pNext = pNode;
	}
	
	pNext = pHead;
	while (pNext)
	{
		cout << pNext->val << endl;
		pNext = pNext->next;
	}
	int index;
	cin >> index;
	Solution s;
	ListNode *pFind = s.FindKthToTail(pHead, index);
	if (pFind == NULL)
	{
		cout << "Not find" << endl;
	}
	else
	{
		cout << "find val = " << pFind->val << endl;
	}
	return 0;
}
