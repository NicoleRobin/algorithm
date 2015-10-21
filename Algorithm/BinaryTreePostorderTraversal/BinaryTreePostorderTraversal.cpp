#include <iostream>
#include <vector>
#include <stack>
using namespace std;

/*
���⣺
�����������������������ֵ��
˼·��
����ջʵ�ֺ������������ݵ�ʱ����Ҫ�б�Ǳ�־�Ѿ����ʹ���Ԫ�أ��������ѭ����
*/

// Definition for a binary tree node.
struct TreeNode {
	int val;    
	TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

struct TrivelTag
{
	TreeNode *pNode;
	bool bTriveled;
};

class Solution {
public:
	vector<int> postorderTraversal(TreeNode* root) {
		vector<int> vecRet;
		stack<TrivelTag> stackTrivel;
		if (root == NULL)
		{
			return vecRet;
		}

		TrivelTag tag;
		tag.pNode = root;
		tag.bTriveled = false;
		stackTrivel.push(tag);
		while (!stackTrivel.empty())
		{
			TreeNode *pNode = stackTrivel.top().pNode;
			if (stackTrivel.top().bTriveled || (pNode->right == NULL && pNode->left == NULL))
			{ // leaf node
				vecRet.push_back(pNode->val);
				stackTrivel.pop();
			}
			else
			{
				stackTrivel.top().bTriveled = true;
				if (pNode->right)
				{
					tag.bTriveled = false;
					tag.pNode = pNode->right;
					stackTrivel.push(tag);
				}
				if (pNode->left)
				{
					tag.bTriveled = false;
					tag.pNode = pNode->left;
					stackTrivel.push(tag);
				}
			}
		}
		return vecRet;
	}
};

int main(void)
{
	return 0;
}