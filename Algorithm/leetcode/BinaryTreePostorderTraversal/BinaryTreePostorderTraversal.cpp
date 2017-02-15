#include <iostream>
#include <vector>
#include <stack>
using namespace std;

/*
题意：
后序遍历二叉树，并给出其值。
思路：
利用栈实现后续遍历，回溯的时候需要有标记标志已经访问过的元素，否则会死循环。
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