#include <iostream>
#include <vector>
using namespace std;

struct TreeNode {
	int val;
	struct TreeNode *left;
	struct TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {
	}
};
/*
* 思路：先序遍历二叉树，找到每一个叶子节点，判断路径和是否为指定的数，如果是，就记录这条路径
*/
class Solution {
public:
	vector<vector<int> > FindPath(TreeNode* root, int expectNumber) {
		if (root == NULL) {
			return vector<vector<int> >();
		}

		vector<vector<int> > vecRet;
		vector<int> vecPath;
		FindPath(root, expectNumber, vecPath, vecRet);
		return vecRet;
	}

	void FindPath(TreeNode *root, int expectNumber, vector<int> vecPath, vector<vector<int> > &vecRet)
	{
		if (root == NULL)
		{
			return;
		}

		expectNumber -= root->val;
		vecPath.push_back(root->val);
		if (root->left == NULL && root->right == NULL)
		{ // 叶子节点
			if (expectNumber == 0)
			{
				vecRet.push_back(vecPath);
				return;
			}
		}

		FindPath(root->left, expectNumber, vecPath, vecRet);
		FindPath(root->right, expectNumber, vecPath, vecRet);
	}
};

int main(void)
{
	TreeNode *pRoot = new TreeNode(1);
	pRoot->left = new TreeNode(2);
	pRoot->right = new TreeNode(3);
	pRoot->left->left = new TreeNode(4);
	pRoot->left->right = new TreeNode(5);
	pRoot->right->left = new TreeNode(6);
	pRoot->right->right = new TreeNode(7);

	Solution s;
	vector<vector<int> > vecRet = s.FindPath(pRoot, 7);
	for (int i = 0; i < vecRet.size(); ++i)
	{
		for (int j = 0; j < vecRet[i].size(); ++j)
		{
			cout << vecRet[i][j] << ", ";
		}
		cout << endl;
	}
	return 0;
}
