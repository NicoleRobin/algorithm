#include <iostream>

using namespace std;

struct TreeNode {
	struct TreeNode *left;
	struct TreeNode *right;
	int val;
	TreeNode(int x) :
	val(x), left(NULL), right(NULL) {
	}
};

class Solution {
public:
	int TreeDepth(TreeNode* pRoot)
	{
		if (pRoot == NULL) {
			return 0;
		}

		int iLeftDepth = TreeDepth(pRoot->left);
		int iRightDepth = TreeDepth(pRoot->right);

		return 1 + (iLeftDepth > iRightDepth ? iLeftDepth : iRightDepth);
	}
};

int main(void)
{
	TreeNode *pRoot = new TreeNode(1);
	pRoot->left = new TreeNode(2);
	pRoot->right = new TreeNode(3);
	pRoot->left->left = new TreeNode(4);
	pRoot->left->right = new TreeNode(5);
	pRoot->left->right->left = new TreeNode(7);
	pRoot->right->right = new TreeNode(6);

	TreeNode *pRoot2 = new TreeNode(1);
	pRoot2->left = new TreeNode(2);

	Solution s;
	cout << s.TreeDepth(pRoot2) << endl;

	return 0;
}
